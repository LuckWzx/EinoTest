# ReAct Agent 执行逻辑

## 一、核心架构

```
┌─────────────────────────────────────────────────────────┐
│                    ReAct Agent                          │
│                                                         │
│   ┌──────────┐    ┌──────────┐    ┌──────────────────┐ │
│   │   LLM    │◄──►│  Tools   │    │  ToolCallChecker │ │
│   │ (ChatModel)│   │ (Add/Sub │    │  (流式工具检测)   │ │
│   │          │    │  /Analyze)│    └──────────────────┘ │
│   └──────────┘    └──────────┘                          │
│        │               │                                 │
│        ▼               ▼                                 │
│   ┌─────────────────────────────────────────────────┐   │
│   │              ReAct Loop (思考→行动→观察)          │   │
│   └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

## 二、ReAct 循环流程

```mermaid
flowchart TD
    A["用户输入 Message<br/>System: 幼儿园老师 persona<br/>User: 183+192-90 的难易程度和答案"] --> B["LLM 分析并决策<br/>(ChatModel)"]
    B --> C{"需要调用工具?"}
    
    C -->|"否: 直接回答"| G["流式输出最终答案"]
    
    C -->|"是: 生成 ToolCalls"| D["ToolCallChecker<br/>检测流中的工具调用"]
    D --> E["执行工具调用<br/>(支持并行: ExecuteSequentially=false)"]
    E --> F["工具结果返回 LLM<br/>作为新的上下文"]
    F --> B
    
    G --> H["结束"]
```

## 三、本例具体执行时序

```mermaid
sequenceDiagram
    actor User as 用户
    participant Agent as ReAct Agent
    participant LLM as ChatModel (Ark)
    participant Add as AddTool
    participant Sub as SubTool
    participant Analyze as AnalyzeTool
    participant InnerLLM as AnalyzeTool 内部 LLM

    User->>Agent: "183+192-90 的难易程度和答案"
    Agent->>LLM: System: 幼儿园老师<br/>User: 183+192-90 的难易程度和答案
    
    Note over LLM: Round 1: 思考决策<br/>需要计算 183+192 和 183+192-90<br/>还需要分析难度

    LLM-->>Agent: ToolCalls:<br/>1. add(183, 192)<br/>2. analyze("183+192-90")
    
    Note over Agent: ToolCallChecker 检测到工具调用<br/>并行执行 (ExecuteSequentially=false)

    par 并行执行
        Agent->>Add: add(183, 192)
        Add-->>Agent: "375"
    and
        Agent->>Analyze: analyze("183+192-90")
        Analyze->>InnerLLM: 判断难度: "183+192-90"
        InnerLLM-->>Analyze: "中等难度，评分4"
        Analyze-->>Agent: "中等难度，评分4"
    end

    Agent->>LLM: 工具结果:<br/>add → 375<br/>analyze → 中等难度<br/>还需要计算 375 - 90

    Note over LLM: Round 2: 分析结果<br/>还需要执行减法

    LLM-->>Agent: ToolCall:<br/>sub(375, 90)

    Agent->>Sub: sub(375, 90)
    Sub-->>Agent: "285"

    Agent->>LLM: 工具结果:<br/>sub → 285<br/>现在有完整信息了

    Note over LLM: Round 3: 综合回答<br/>难度: 中等<br/>答案: 285<br/>以幼儿园老师口吻回复

    LLM-->>Agent: 最终回答:<br/>"小朋友们，这道题难度中等哦~<br/>183+192-90 的答案是 285！"
    
    Agent-->>User: 流式输出最终回答
```

## 四、关键组件说明

### 4.1 Agent 配置

| 配置项 | 值 | 说明 |
|--------|-----|------|
| `ToolCallingModel` | Ark ChatModel | 驱动 ReAct 循环的 LLM |
| `Tools` | AddTool, SubTool, AnalyzeTool | 可供调用的工具集 |
| `ExecuteSequentially` | `false` | 工具并行执行 |
| `StreamToolCallChecker` | 自定义函数 | 从流中检测是否有工具调用 |

### 4.2 三个工具

| 工具 | 功能 | 输入 | 输出 |
|------|------|------|------|
| AddTool | 加法 | `{a: int, b: int}` | 和 (string) |
| SubTool | 减法 | `{a: int, b: int}` | 差 (string) |
| AnalyzeTool | 难度评估 | `{content: string}` | 难度描述 (调用内部 LLM) |

### 4.3 ToolCallChecker

```
流式读取 LLM 输出 → 逐条检查 msg.ToolCalls → 有则返回 true → 无则继续读 → EOF 返回 false
```

作用：在流式场景下提前判断 LLM 是打算调用工具还是直接回答，避免 Agent 等待完整输出后才做决策。

### 4.4 LoggerCallback

通过 `callbacks` 机制注入，监控 Agent 执行过程：
- **OnStart**: 打印节点输入
- **OnEnd**: 打印节点输出
- **OnEndWithStreamOutput**: 打印 Graph 层级的流式输出
- **OnError**: 打印错误信息

## 五、数据流总结

```
User Input ([]*schema.Message)
    │
    ▼
┌──────────────────────────────────────┐
│         ReAct Agent (Graph)          │
│                                      │
│  ┌────┐   ToolCall?   ┌────────┐   │
│  │LLM │──────────────►│ Tools  │   │
│  │    │◄──────────────│ Node   │   │
│  └────┘   ToolResult  └────────┘   │
│     │                               │
│     │ No ToolCall (直接回答)          │
│     ▼                               │
│  Final Answer                       │
└──────────────────────────────────────┘
    │
    ▼
Stream Output (*schema.StreamReader[*schema.Message])
    │
    ▼
Final Content: "小朋友们，这道题是中等难度，答案是 285 哦~"
```
