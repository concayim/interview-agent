package questions

import (
	"fmt"
	"sort"
	"strings"

	"interview-agent/internal/domain"
)

var bank = []domain.Question{
	{ID: "go-01", Language: "golang", Difficulty: "easy", Prompt: "Go 的 goroutine 和操作系统线程有什么区别？调度器为什么能支撑大量并发？", StandardAnswer: "goroutine 是由 Go runtime 管理的轻量执行单元，初始栈很小且可动态增长。GMP 调度模型用 G 表示 goroutine、M 表示线程、P 表示执行 Go 代码所需的调度资源，通过 work stealing、网络轮询和抢占把大量 G 复用到少量 M 上，因此创建和切换成本通常低于线程。", KeyPoints: []string{"goroutine", "gmp", "动态栈/stack", "work stealing/工作窃取", "复用/m:n"}, Tags: []string{"Go", "高并发", "调度器"}},
	{ID: "go-02", Language: "golang", Difficulty: "easy", Prompt: "channel 的关闭原则是什么？读取已关闭的 channel 会发生什么？", StandardAnswer: "通常由发送方关闭 channel，接收方不应关闭，且应避免重复关闭。关闭后不能继续发送，否则会 panic；接收方仍可读完缓冲区，之后立即得到元素类型零值，双返回值读取中的 ok 为 false。nil channel 的收发会永久阻塞。", KeyPoints: []string{"发送方关闭/sender", "panic", "缓冲/buffer", "零值/zero value", "ok false"}, Tags: []string{"Go", "channel", "并发"}},
	{ID: "go-03", Language: "golang", Difficulty: "medium", Prompt: "请解释 Go interface 的底层语义，以及为什么一个包含 nil 指针的 interface 不等于 nil。", StandardAnswer: "interface 值可理解为动态类型与动态值二元组。只有动态类型和动态值都为空时，interface 才等于 nil。把一个值为 nil 的 *T 赋给 interface 后，动态类型仍是 *T，所以该 interface 不为 nil；调用方法时还需留意接收者是否能处理 nil。", KeyPoints: []string{"动态类型/dynamic type", "动态值/dynamic value", "二元组/pair", "*t", "都为空/both nil"}, Tags: []string{"Go", "interface", "类型系统"}},
	{ID: "go-04", Language: "golang", Difficulty: "medium", Prompt: "怎样系统地定位 Go 服务的内存上涨？请说出你的排查路径和常用工具。", StandardAnswer: "先区分堆增长、goroutine 泄漏、缓存增长与进程外内存，再结合指标确认趋势。使用 net/http/pprof 或 runtime/pprof 抓取 heap、allocs、goroutine profile，用 go tool pprof 的 top/list/diff 分析分配热点与对象保留；必要时结合 trace、GC 指标和逃逸分析。修复后用同流量压测与 profile 对比验证。", KeyPoints: []string{"pprof", "heap", "goroutine", "top/list/diff", "gc", "对比/verify"}, Tags: []string{"Go", "性能优化", "pprof"}},
	{ID: "go-05", Language: "golang", Difficulty: "hard", Prompt: "设计一个可优雅停止的 Go worker pool。你会如何处理背压、取消、panic 和任务结果？", StandardAnswer: "用有界任务 channel 控制队列并形成背压；context 贯穿提交与 worker，停止时先拒绝新任务，再取消或关闭任务源，并用 WaitGroup 等待在途任务。每个 worker 在任务边界 recover 并记录堆栈，结果通过独立有界 channel 或回调返回，发送结果也要响应取消。需要明确 drain 还是立即终止语义，并避免发送方与关闭方竞态。", KeyPoints: []string{"有界/bounded", "背压/backpressure", "context", "waitgroup", "recover", "drain"}, Tags: []string{"Go", "高并发", "架构设计"}},
	{ID: "go-06", Language: "golang", Difficulty: "hard", Prompt: "Go map 并发读写为什么不安全？生产代码中你会如何在 mutex、sync.Map 和分片 map 之间选择？", StandardAnswer: "普通 map 的桶扩容与写入不是并发安全的，并发读写会产生数据竞争，运行时也可能直接报错。读写模式稳定且需要类型安全时优先 map 加 RWMutex；键集合稳定、读多写少或不同 goroutine 操作不同键时可考虑 sync.Map；热点高、锁竞争明显且可按键散列时用分片 map。选择应以 race detector 与 benchmark/profile 为依据。", KeyPoints: []string{"数据竞争/data race", "扩容/resize", "rwmutex", "sync.map", "分片/shard", "benchmark"}, Tags: []string{"Go", "map", "高并发"}},

	{ID: "java-01", Language: "java", Difficulty: "easy", Prompt: "HashMap 在 Java 8 中如何处理哈希冲突？扩容时发生什么？", StandardAnswer: "HashMap 使用数组加链表/红黑树。键先经扰动计算桶位置，冲突节点进入同一桶；链表长度达到阈值且容量足够时树化，过短时可退化。容量通常按 2 倍扩容，节点依据旧容量对应的哈希位决定留在原索引或移动 oldCap 的偏移位置。HashMap 本身非线程安全。", KeyPoints: []string{"数组/array", "链表/linked list", "红黑树/red-black", "树化/treeify", "2倍/double", "线程不安全/not thread-safe"}, Tags: []string{"Java", "HashMap", "数据结构"}},
	{ID: "java-02", Language: "java", Difficulty: "easy", Prompt: "== 与 equals 有什么区别？实现 equals 时为什么通常也要实现 hashCode？", StandardAnswer: "对引用类型，== 比较是否指向同一对象，equals 表达逻辑相等且默认实现与 == 一致。若两个对象 equals 为 true，hashCode 必须相同，否则基于哈希的集合无法正确查找；反向不要求成立。实现时还应满足自反、对称、传递、一致与非空性。", KeyPoints: []string{"引用/reference", "逻辑相等/logical", "hashcode", "哈希集合/hash", "相同/same"}, Tags: []string{"Java", "对象模型", "集合"}},
	{ID: "java-03", Language: "java", Difficulty: "medium", Prompt: "解释 JVM 中 volatile 的可见性与有序性保证，它为什么不能替代所有锁？", StandardAnswer: "volatile 写与后续对同一变量的读之间建立 happens-before，保证刷新/读取主内存，并通过内存屏障限制相关指令重排。但复合操作如 i++ 仍包含读改写多个步骤，不具备原子性；需要原子类、锁或其他同步机制维护跨变量不变式。", KeyPoints: []string{"happens-before", "可见性/visibility", "有序性/ordering", "内存屏障/memory barrier", "不保证原子性/not atomic"}, Tags: []string{"Java", "JVM", "并发"}},
	{ID: "java-04", Language: "java", Difficulty: "medium", Prompt: "Spring 事务在哪些常见场景会失效？你会怎样验证事务边界？", StandardAnswer: "常见问题包括同类方法自调用绕过代理、方法非 public、异常被吞掉、默认只对未检查异常回滚、传播级别不合预期、切错数据源或对象不是 Spring Bean。应检查代理是否生效、事务管理器与传播配置，用集成测试制造异常并验证数据库状态，同时观察事务日志。", KeyPoints: []string{"代理/proxy", "自调用/self-invocation", "异常/exception", "传播/propagation", "事务管理器/transaction manager", "集成测试"}, Tags: []string{"Java", "Spring Boot", "事务"}},
	{ID: "java-05", Language: "java", Difficulty: "hard", Prompt: "线上出现 Full GC 频繁，你会如何定位？不要只列工具，请说明判断顺序。", StandardAnswer: "先从 GC 日志与监控确认收集器、停顿、晋升速率、老年代占用和回收效果，判断是分配过快、内存泄漏、元空间、直接内存还是堆设置问题。安全地获取 class histogram 与 heap dump，用 MAT/YourKit 等看 dominator tree 和 GC roots；结合 JFR、代码发布与流量变化定位分配源。修复或调参后用同负载对比验证，避免把增大堆当作唯一方案。", KeyPoints: []string{"gc日志/gc log", "老年代/old generation", "heap dump", "gc roots", "dominator", "jfr", "验证/compare"}, Tags: []string{"Java", "JVM", "性能优化"}},
	{ID: "java-06", Language: "java", Difficulty: "hard", Prompt: "如何设计一个幂等的订单创建接口，同时应对客户端重试和消息重复消费？", StandardAnswer: "由客户端或服务端生成业务唯一的幂等键，数据库以唯一约束作为最终防线；在同一事务内记录幂等请求与订单结果，重复请求返回已存在结果。消息消费端也以业务键做去重或状态机条件更新，结合 outbox/本地消息表保证数据库变更与事件发布的一致性。需要定义幂等记录过期、进行中请求和失败重试语义。", KeyPoints: []string{"幂等键/idempotency key", "唯一约束/unique", "事务/transaction", "去重/dedup", "outbox", "状态机/state machine"}, Tags: []string{"Java", "微服务", "分布式系统"}},

	{ID: "python-01", Language: "python", Difficulty: "easy", Prompt: "Python 的可变默认参数为什么危险？应该如何改写？", StandardAnswer: "默认参数在函数定义时只求值一次，因此列表或字典等可变对象会被多次调用共享，导致状态意外累积。通常将默认值设为 None，在函数体内创建新对象；若 None 本身是合法值，可使用独立哨兵对象。", KeyPoints: []string{"定义时/definition time", "只求值一次/once", "共享/shared", "none", "哨兵/sentinel"}, Tags: []string{"Python", "函数", "语言基础"}},
	{ID: "python-02", Language: "python", Difficulty: "easy", Prompt: "list、tuple、set 和 dict 的核心差异与典型使用场景是什么？", StandardAnswer: "list 有序可变，适合序列；tuple 有序不可变，元素可哈希时可作为键；set 保存唯一的可哈希元素，适合集合运算和成员判断；dict 保存键值映射，键需可哈希。现代 Python 的 dict 保留插入顺序，但这不等同于按键排序。", KeyPoints: []string{"list", "tuple", "set", "dict", "可哈希/hashable", "插入顺序/insertion order"}, Tags: []string{"Python", "数据结构", "集合"}},
	{ID: "python-03", Language: "python", Difficulty: "medium", Prompt: "请解释 GIL。它对 CPU 密集与 I/O 密集任务分别意味着什么？", StandardAnswer: "在常见 CPython 中，GIL 使同一进程同一时刻通常只有一个线程执行 Python 字节码。I/O 阻塞时会释放 GIL，因此线程仍适合许多 I/O 密集任务；纯 Python CPU 密集任务通常不能靠多线程线性加速，可用多进程、原生扩展或其他运行时。GIL 不等于业务代码天然线程安全。", KeyPoints: []string{"cpython", "字节码/bytecode", "io", "多进程/multiprocessing", "线程安全/thread-safe"}, Tags: []string{"Python", "GIL", "并发"}},
	{ID: "python-04", Language: "python", Difficulty: "medium", Prompt: "asyncio 的事件循环如何工作？什么行为会把整个异步服务卡住？", StandardAnswer: "事件循环调度协程，协程在 await 未就绪的 I/O 或 Future 时主动让出控制权，完成后再恢复。协程中直接执行长时间 CPU 计算、同步网络/文件调用、time.sleep 或阻塞锁会卡住事件循环；应改用异步库，或通过线程池/进程池卸载阻塞和 CPU 工作，并设置超时与取消处理。", KeyPoints: []string{"事件循环/event loop", "await", "主动让出/yield", "阻塞/blocking", "线程池/thread pool", "超时/timeout"}, Tags: []string{"Python", "asyncio", "高并发"}},
	{ID: "python-05", Language: "python", Difficulty: "hard", Prompt: "一个 Python API 服务延迟偶发升高，你会如何区分是代码、数据库、GC 还是下游依赖导致？", StandardAnswer: "先用端到端指标和分布式追踪按请求拆分时间，关联吞吐、错误率和资源饱和度。检查 event loop lag/线程池队列、慢 SQL 与连接池、下游调用分位数及超时重试，再用 py-spy/cProfile 等采样 CPU，用 tracemalloc 与 GC 指标观察分配和回收。以时间线和 trace 证据确定瓶颈，做单变量调整并回归验证。", KeyPoints: []string{"分布式追踪/tracing", "慢sql/slow sql", "连接池/connection pool", "下游/downstream", "py-spy", "gc", "验证/verify"}, Tags: []string{"Python", "性能优化", "可观测性"}},
	{ID: "python-06", Language: "python", Difficulty: "hard", Prompt: "如何让一个消费 Kafka 的 Python 服务做到至少一次投递下的业务幂等？", StandardAnswer: "不要仅依赖消息 offset 去重，应选取稳定业务键，在数据库用唯一约束或带版本的状态更新保证重复处理无副作用。业务写入成功后再提交 offset；若还要发事件，可在同一数据库事务写 outbox，再异步发布。明确 poison message 的重试与死信策略，并监控消费延迟和重复率。", KeyPoints: []string{"业务键/business key", "唯一约束/unique", "提交offset/commit", "outbox", "死信/dead letter", "重试/retry"}, Tags: []string{"Python", "Kafka", "分布式系统"}},

	{ID: "cpp-01", Language: "cpp", Difficulty: "easy", Prompt: "指针和引用的主要区别是什么？什么场景下你更倾向使用引用？", StandardAnswer: "指针是对象，可为空、可重新指向并支持指针运算；引用通常必须在初始化时绑定且语义上不可改绑，使用方式更像对象别名。表达必需存在的输入参数或运算符返回值时常用引用；需要可选性、所有权表达或数组遍历时使用指针或智能指针。", KeyPoints: []string{"可为空/null", "重新指向/reseat", "初始化/initialization", "别名/alias", "所有权/ownership"}, Tags: []string{"C++", "指针", "引用"}},
	{ID: "cpp-02", Language: "cpp", Difficulty: "easy", Prompt: "什么是 RAII？它如何帮助避免资源泄漏？", StandardAnswer: "RAII 把资源生命周期绑定到对象生命周期：构造函数获得资源，析构函数释放资源。离开作用域时无论正常返回还是异常展开，析构都会执行，从而可靠释放内存、文件、锁等资源。标准容器、unique_ptr、lock_guard 都是典型实践。", KeyPoints: []string{"对象生命周期/object lifetime", "构造/constructor", "析构/destructor", "异常/exception", "unique_ptr", "lock_guard"}, Tags: []string{"C++", "RAII", "资源管理"}},
	{ID: "cpp-03", Language: "cpp", Difficulty: "medium", Prompt: "unique_ptr、shared_ptr、weak_ptr 分别表达什么所有权？shared_ptr 有什么典型陷阱？", StandardAnswer: "unique_ptr 表达独占所有权并支持移动；shared_ptr 用引用计数表达共享所有权；weak_ptr 不增加强引用计数，用于观察 shared_ptr 管理的对象并打破环。shared_ptr 的陷阱包括循环引用、控制块额外开销、误从同一裸指针创建多个控制块，以及共享所有权掩盖生命周期设计。", KeyPoints: []string{"独占/unique", "引用计数/reference count", "weak_ptr", "循环引用/cycle", "控制块/control block", "裸指针/raw pointer"}, Tags: []string{"C++", "智能指针", "内存管理"}},
	{ID: "cpp-04", Language: "cpp", Difficulty: "medium", Prompt: "move semantics 解决了什么问题？std::move 本身会移动对象吗？", StandardAnswer: "移动语义允许资源拥有者把内部资源转移给新对象，避免昂贵深拷贝。std::move 本身只是把表达式转换为右值引用，真正的转移发生在被调用的移动构造或移动赋值中。被移动对象必须保持有效但状态通常未指定，仍需满足析构和重新赋值要求。", KeyPoints: []string{"资源转移/transfer", "避免深拷贝/avoid copy", "右值引用/rvalue", "转换/cast", "有效但未指定/valid but unspecified"}, Tags: []string{"C++", "移动语义", "性能优化"}},
	{ID: "cpp-05", Language: "cpp", Difficulty: "hard", Prompt: "怎样实现一个线程安全的有界队列？请说明条件变量的等待谓词和关闭语义。", StandardAnswer: "用 mutex 保护队列、容量和关闭标记，用 not_empty/not_full 条件变量分别唤醒消费者与生产者。wait 必须带谓词并在锁内重新检查，以应对虚假唤醒。push 在满时等待或按策略超时，pop 在空时等待；关闭后拒绝 push，可选择让消费者 drain 剩余元素，全部取完后 pop 返回关闭状态。状态变更后在合适时机 notify，并保证对象销毁前线程退出。", KeyPoints: []string{"mutex", "条件变量/condition variable", "谓词/predicate", "虚假唤醒/spurious", "有界/bounded", "关闭/close", "drain"}, Tags: []string{"C++", "并发", "数据结构"}},
	{ID: "cpp-06", Language: "cpp", Difficulty: "hard", Prompt: "线上 C++ 服务出现偶发 use-after-free，你会如何定位和修复？", StandardAnswer: "优先在可复现环境启用 AddressSanitizer，保留符号并收集完整栈；并发相关问题配合 ThreadSanitizer，但通常分开运行。审计所有权与跨线程生命周期，减少裸 owning pointer，使用 RAII 和合适智能指针；注意回调捕获 this、容器迭代器失效和异步任务。修复后添加压力/竞态测试并持续在 CI 跑 sanitizer。", KeyPoints: []string{"addresssanitizer/asan", "threadsanitizer/tsan", "符号/symbol", "所有权/ownership", "this", "迭代器/iterator", "ci"}, Tags: []string{"C++", "内存安全", "调试"}},
}

var foundationBank = []domain.Question{
	{ID: "foundation-01", Language: "foundation", Difficulty: "easy", Prompt: "进程和线程的核心区别是什么？一次线程切换通常需要保存哪些状态？", StandardAnswer: "进程是资源分配与隔离的基本单位，拥有独立虚拟地址空间等资源；线程是进程内的执行单元，共享进程资源但拥有自己的栈、寄存器和调度状态。线程切换通常保存和恢复程序计数器、寄存器、栈指针及调度上下文；跨进程切换还可能涉及地址空间和页表相关开销。", KeyPoints: []string{"资源隔离/process isolation", "共享地址空间/shared memory", "栈/stack", "寄存器/register", "调度上下文/context switch"}, Tags: []string{"操作系统", "进程", "线程"}},
	{ID: "foundation-02", Language: "foundation", Difficulty: "easy", Prompt: "TCP 为什么需要三次握手，而不是两次？", StandardAnswer: "三次握手让双方确认彼此的发送和接收能力，并同步初始序列号。仅两次无法让服务端确认客户端已经收到服务端的序列号与确认，旧的延迟连接请求也更容易造成半开或错误建连。第三次 ACK 完成双向状态确认。", KeyPoints: []string{"双向/bidirectional", "序列号/sequence number", "ack", "旧连接/delayed request", "状态确认/state"}, Tags: []string{"计算机网络", "TCP", "基础"}},
	{ID: "foundation-03", Language: "foundation", Difficulty: "medium", Prompt: "数据库索引为什么常用 B+ 树？它相对哈希索引和二叉树有什么优势？", StandardAnswer: "B+ 树分支多、树高低，节点大小适合页式存储，能减少磁盘或缓存页访问；数据集中在叶子并通过链表连接，适合范围查询和顺序扫描。哈希索引适合等值但不擅长范围与排序，普通二叉树分支少、树高更高且局部性较差。", KeyPoints: []string{"树高低/height", "页/page", "叶子链表/leaf", "范围查询/range", "哈希等值/hash equality"}, Tags: []string{"数据库", "B+树", "索引"}},
	{ID: "foundation-04", Language: "foundation", Difficulty: "medium", Prompt: "什么是缓存穿透、击穿和雪崩？分别如何治理？", StandardAnswer: "穿透是查询不存在的数据绕过缓存，可用参数校验、空值缓存或布隆过滤器；击穿是热点键失效瞬间大量回源，可用互斥重建、逻辑过期或预热；雪崩是大量键同时失效或缓存整体不可用，可打散过期时间、多级缓存、限流降级并提高缓存集群可用性。", KeyPoints: []string{"穿透/penetration", "布隆过滤器/bloom", "击穿/hot key", "互斥/mutex", "雪崩/avalanche", "过期打散/jitter"}, Tags: []string{"缓存", "Redis", "高并发"}},
	{ID: "foundation-05", Language: "foundation", Difficulty: "hard", Prompt: "在分布式系统中，如何理解一致性、可用性和分区容错之间的取舍？", StandardAnswer: "网络分区发生时，系统必须在强一致性与持续可用之间做选择，这就是 CAP 的关键语境，而不是日常无分区时只能三选二。CP 系统可能拒绝或延迟部分请求以保持一致，AP 系统继续服务但允许临时不一致并通过冲突解决最终收敛。实际设计还需结合一致性级别、业务不变量、延迟和故障模型。", KeyPoints: []string{"网络分区/partition", "一致性/consistency", "可用性/availability", "cp", "ap", "业务不变量/invariant"}, Tags: []string{"分布式系统", "CAP", "架构设计"}},
	{ID: "foundation-06", Language: "foundation", Difficulty: "hard", Prompt: "设计一个限流器时，令牌桶和漏桶有什么差异？分布式部署还需考虑什么？", StandardAnswer: "令牌桶按速率补充令牌，允许在桶容量范围内突发；漏桶以较稳定速率流出，更强调平滑。分布式部署需明确全局还是单实例配额，处理时钟、原子扣减、热点、网络失败和配置动态下发；可用 Redis/Lua 或专用限流服务，但要设计失败时放行还是拒绝以及本地兜底。", KeyPoints: []string{"令牌桶/token bucket", "突发/burst", "漏桶/leaky bucket", "平滑/smooth", "原子/atomic", "失败策略/failure"}, Tags: []string{"算法", "限流", "分布式系统"}},
}

func Languages() []string { return []string{"golang", "java", "python", "cpp"} }

func All() []domain.Question {
	result := make([]domain.Question, 0, len(bank)+len(foundationBank))
	result = append(result, bank...)
	result = append(result, foundationBank...)
	return result
}

func Select(language, difficulty string, count int, resumeKeywords []string) ([]domain.Question, error) {
	return selectFrom(bank, language, difficulty, count, resumeKeywords)
}

func SelectWithFoundation(language, difficulty string, count int, resumeKeywords []string, includeFoundation bool) ([]domain.Question, error) {
	if !includeFoundation {
		return Select(language, difficulty, count, resumeKeywords)
	}
	if count < 1 {
		count = 5
	}
	languageCandidates, err := selectFrom(bank, language, difficulty, 20, resumeKeywords)
	if err != nil {
		return nil, err
	}
	foundationCandidates, err := selectFrom(foundationBank, "foundation", difficulty, 20, resumeKeywords)
	if err != nil {
		return languageCandidates[:min(count, len(languageCandidates))], nil
	}
	target := min(count, len(languageCandidates)+len(foundationCandidates))
	foundationCount := max(1, target/3)
	if target-foundationCount > len(languageCandidates) {
		foundationCount = target - len(languageCandidates)
	}
	foundationCount = min(foundationCount, len(foundationCandidates))
	languageCount := min(target-foundationCount, len(languageCandidates))
	if languageCount+foundationCount < target {
		foundationCount = min(len(foundationCandidates), target-languageCount)
	}
	languageQuestions := languageCandidates[:languageCount]
	foundationQuestions := foundationCandidates[:foundationCount]
	result := make([]domain.Question, 0, target)
	for languageIndex, foundationIndex := 0, 0; len(result) < target; {
		for added := 0; added < 2 && languageIndex < len(languageQuestions); added++ {
			result = append(result, languageQuestions[languageIndex])
			languageIndex++
		}
		if foundationIndex < len(foundationQuestions) {
			result = append(result, foundationQuestions[foundationIndex])
			foundationIndex++
		}
		if languageIndex >= len(languageQuestions) {
			for foundationIndex < len(foundationQuestions) && len(result) < target {
				result = append(result, foundationQuestions[foundationIndex])
				foundationIndex++
			}
		}
	}
	return result, nil
}

func selectFrom(source []domain.Question, language, difficulty string, count int, resumeKeywords []string) ([]domain.Question, error) {
	if count < 1 {
		count = 5
	}
	if count > 20 {
		count = 20
	}
	candidates := make([]domain.Question, 0)
	for _, q := range source {
		if q.Language == language && (difficulty == "mixed" || q.Difficulty == difficulty) {
			candidates = append(candidates, q)
		}
	}
	if len(candidates) == 0 {
		return nil, fmt.Errorf("没有找到语言 %q、难度 %q 的题目", language, difficulty)
	}
	keywords := make(map[string]bool, len(resumeKeywords))
	for _, keyword := range resumeKeywords {
		keywords[strings.ToLower(keyword)] = true
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		return relevance(candidates[i], keywords) > relevance(candidates[j], keywords)
	})
	if count > len(candidates) {
		count = len(candidates)
	}
	return append([]domain.Question(nil), candidates[:count]...), nil
}

func relevance(q domain.Question, keywords map[string]bool) int {
	score := 0
	for _, tag := range q.Tags {
		if keywords[strings.ToLower(tag)] {
			score++
		}
	}
	return score
}
