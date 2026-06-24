import { useEffect, useRef, useState } from 'react'
import type { CSSProperties, Dispatch, SetStateAction } from 'react'
import {
  ArrowLeft,
  ArrowRight,
  BookOpen,
  Bookmark,
  BookmarkCheck,
  Bot,
  BrainCircuit,
  Check,
  ChevronDown,
  ChevronRight,
  CircleAlert,
  Clock3,
  Code2,
  FileCheck2,
  FileText,
  Globe2,
  Gauge,
  Lightbulb,
  LoaderCircle,
  MessageSquareText,
  Newspaper,
  PlayCircle,
  RefreshCw,
  Send,
  Settings2,
  ShieldCheck,
  Sparkles,
  Target,
  LibraryBig,
  UploadCloud,
  UserRound,
  X,
  Zap,
} from 'lucide-react'
import { api } from './api'
import type { Difficulty, Evaluation, KnowledgeBase, Language, LearningResource, ModelConfig, QAItem, Question, Report, Resume, Session, SkillCatalog } from './types'

type Screen = 'setup' | 'interview' | 'review' | 'learning' | 'knowledge'
type Turn = { question: Question; answer: string; evaluation: Evaluation; mainIndex: number }
type Toast = { type: 'success' | 'error'; message: string }

const languageOptions: { value: Language; label: string; short: string; caption: string; color: string }[] = [
  { value: 'golang', label: 'Golang', short: 'Go', caption: '并发 · Runtime · 工程化', color: '#6bd7e8' },
  { value: 'java', label: 'Java', short: 'J', caption: 'JVM · Spring · 分布式', color: '#ffb46e' },
  { value: 'python', label: 'Python', short: 'Py', caption: '语言 · 异步 · 服务端', color: '#e8d875' },
  { value: 'cpp', label: 'C++', short: 'C++', caption: '内存 · 并发 · 性能', color: '#ac9cff' },
]

const difficultyOptions: { value: Difficulty; label: string; caption: string }[] = [
  { value: 'mixed', label: '智能混合', caption: '从基础逐步深入' },
  { value: 'easy', label: '基础', caption: '2 道专项题' },
  { value: 'medium', label: '进阶', caption: '2 道专项题' },
  { value: 'hard', label: '挑战', caption: '2 道专项题' },
]

function App() {
  const [screen, setScreen] = useState<Screen>('setup')
  const [settingsOpen, setSettingsOpen] = useState(false)
  const [resume, setResume] = useState<Resume>()
  const [session, setSession] = useState<Session>()
  const [report, setReport] = useState<Report>()
  const [turns, setTurns] = useState<Turn[]>([])
  const [modelConfig, setModelConfig] = useState<ModelConfig>({ baseUrl: '', model: '', enabled: false, hasApiKey: false })
  const [catalog, setCatalog] = useState<SkillCatalog>()
  const [catalogLoading, setCatalogLoading] = useState(true)
  const [catalogError, setCatalogError] = useState('')
  const [toast, setToast] = useState<Toast>()

  const loadCatalog = async () => {
    setCatalogLoading(true)
    setCatalogError('')
    try { setCatalog(await api.skills()) }
    catch (error) { setCatalogError(error instanceof Error ? error.message : 'Skill 目录加载失败') }
    finally { setCatalogLoading(false) }
  }

  useEffect(() => {
    void loadCatalog()
    api.getModelConfig().then(setModelConfig).catch((error) => setToast({ type: 'error', message: `模型配置加载失败：${error.message}` }))
  }, [])
  useEffect(() => {
    if (!toast) return
    const timer = window.setTimeout(() => setToast(undefined), 3200)
    return () => window.clearTimeout(timer)
  }, [toast])
  useEffect(() => {
    document.querySelector('.app-main')?.scrollTo({ top: 0, behavior: 'instant' })
  }, [screen])

  const start = (next: Session) => {
    setSession(next)
    setTurns([])
    setReport(undefined)
    setScreen('interview')
  }

  const finish = (nextReport: Report) => {
    setReport(nextReport)
    setScreen('review')
  }

  const restart = () => {
    setSession(undefined)
    setReport(undefined)
    setTurns([])
    setScreen('setup')
  }

  return (
    <div className="app-shell">
      <Sidebar screen={screen} modelReady={modelConfig.enabled && modelConfig.hasApiKey} onHome={restart} onLearning={() => setScreen('learning')} onKnowledge={() => setScreen('knowledge')} onSettings={() => setSettingsOpen(true)} />
      <main className="app-main">
        {screen === 'setup' && <Setup catalog={catalog} catalogLoading={catalogLoading} catalogError={catalogError} onRetryCatalog={loadCatalog} resume={resume} onResume={setResume} onStart={start} modelReady={modelConfig.enabled && modelConfig.hasApiKey} notify={setToast} />}
        {screen === 'interview' && session && (
          <InterviewScreen session={session} turns={turns} setTurns={setTurns} setSession={setSession} onFinish={finish} onBack={restart} notify={setToast} />
        )}
        {screen === 'review' && report && <ReviewScreen report={report} onRestart={restart} />}
        {screen === 'learning' && <LearningPage catalog={catalog} notify={setToast} />}
        {screen === 'knowledge' && <KnowledgePage notify={setToast} />}
      </main>
      <SettingsDrawer open={settingsOpen} value={modelConfig} onClose={() => setSettingsOpen(false)} onSaved={setModelConfig} notify={setToast} />
      {toast && <div className={`toast toast-${toast.type}`}>{toast.type === 'success' ? <Check size={18} /> : <CircleAlert size={18} />}{toast.message}</div>}
    </div>
  )
}

function Sidebar({ screen, modelReady, onHome, onLearning, onKnowledge, onSettings }: { screen: Screen; modelReady: boolean; onHome: () => void; onLearning: () => void; onKnowledge: () => void; onSettings: () => void }) {
  return (
    <aside className="sidebar">
      <button className="brand" onClick={onHome} aria-label="返回首页">
        <span className="brand-mark"><Sparkles size={20} /></span>
        <span><strong>Interview</strong><small>Copilot</small></span>
      </button>
      <nav className="step-nav" aria-label="面试流程">
        <div className={`nav-step ${screen === 'setup' ? 'active' : ''}`}><span>01</span><div><strong>准备</strong><small>简历与方向</small></div></div>
        <div className={`nav-line ${screen !== 'setup' ? 'filled' : ''}`} />
        <div className={`nav-step ${screen === 'interview' ? 'active' : screen === 'review' ? 'done' : ''}`}><span>{screen === 'review' ? <Check size={14} /> : '02'}</span><div><strong>面试</strong><small>沉浸式问答</small></div></div>
        <div className={`nav-line ${screen === 'review' ? 'filled' : ''}`} />
        <div className={`nav-step ${screen === 'review' ? 'active' : ''}`}><span>03</span><div><strong>复盘</strong><small>答案与建议</small></div></div>
      </nav>
      <div className="sidebar-spacer" />
      <nav className="explore-nav" aria-label="学习与知识">
        <button className={screen === 'learning' ? 'active' : ''} onClick={onLearning}><BookOpen size={17} /><span><strong>学习中心</strong><small>权威文章与视频</small></span></button>
        <button className={screen === 'knowledge' ? 'active' : ''} onClick={onKnowledge}><LibraryBig size={17} /><span><strong>知识库</strong><small>分库 QA 检索</small></span></button>
      </nav>
      <div className={`model-pill ${modelReady ? 'online' : ''}`}><span className="status-dot" /><div><strong>{modelReady ? 'AI 增强已开启' : '本地模式'}</strong><small>{modelReady ? 'Eino 模型在线' : '规则评分可用'}</small></div></div>
      <button className="sidebar-button" onClick={onSettings}><Settings2 size={17} />模型设置</button>
      <div className="privacy-note"><ShieldCheck size={15} /><span>简历仅在本机解析</span></div>
    </aside>
  )
}

function Setup({ catalog, catalogLoading, catalogError, onRetryCatalog, resume, onResume, onStart, modelReady, notify }: { catalog?: SkillCatalog; catalogLoading: boolean; catalogError: string; onRetryCatalog: () => Promise<void>; resume?: Resume; onResume: (value: Resume) => void; onStart: (value: Session) => void; modelReady: boolean; notify: (value: Toast) => void }) {
  const [candidateName, setCandidateName] = useState('')
  const [industry, setIndustry] = useState('computer')
  const [domainSkillId, setDomainSkillId] = useState('computer-golang')
  const [interviewerSkillId, setInterviewerSkillId] = useState('echo-coach')
  const [includeFoundation, setIncludeFoundation] = useState(true)
  const [difficulty, setDifficulty] = useState<Difficulty>('mixed')
  const [questionCount, setQuestionCount] = useState(5)
  const [uploading, setUploading] = useState(false)
  const [starting, setStarting] = useState(false)
  const [dragging, setDragging] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null)

  const maxQuestionCount = difficulty === 'mixed' ? (includeFoundation ? 12 : 6) : (includeFoundation ? 4 : 2)
  useEffect(() => { if (questionCount > maxQuestionCount) setQuestionCount(maxQuestionCount) }, [difficulty, includeFoundation, maxQuestionCount, questionCount])

  const upload = async (file?: File) => {
    if (!file) return
    if (!/\.(pdf|docx)$/i.test(file.name)) { notify({ type: 'error', message: '请选择 PDF 或 DOCX 简历' }); return }
    if (file.size > 10 * 1024 * 1024) { notify({ type: 'error', message: '附件不能超过 10 MB' }); return }
    setUploading(true)
    try { onResume(await api.uploadResume(file)); notify({ type: 'success', message: '简历解析完成，已提取技术关键词' }) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '简历解析失败' }) }
    finally { setUploading(false) }
  }

  const start = async () => {
    setStarting(true)
    try { onStart(await api.startInterview({ candidateName, resumeId: resume?.id, domainSkillId, interviewerSkillId, includeFoundation, difficulty, questionCount })) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '面试创建失败' }) }
    finally { setStarting(false) }
  }

  const domainSkills = (catalog?.domains ?? []).filter((skill) => skill.industry === industry && skill.language !== 'foundation')
  const selectedInterviewer = catalog?.interviewers.find((skill) => skill.id === interviewerSkillId)
  const followUpRounds = selectedInterviewer?.followUpRounds ?? 2

  return (
    <div className="setup-page page-enter">
      <header className="hero">
        <div><span className="eyebrow"><Sparkles size={14} /> AI 技术面试陪练</span><h1>把每一次回答，<br /><em>练成底气。</em></h1><p>上传简历，选择技术方向。面试官会围绕你的经历发问，并在最后给出有标准答案的完整复盘。</p></div>
        <div className="hero-orbit" aria-hidden="true"><span className="orbit orbit-one" /><span className="orbit orbit-two" /><BrainCircuit size={56} /></div>
      </header>

      {catalogError && <div className="skill-load-alert"><CircleAlert size={19} /><div><strong>Skill 目录没有加载成功</strong><p>{catalogError}。请确认本地服务已启动后重试。</p></div><button onClick={() => void onRetryCatalog()}><RefreshCw size={15} />重新加载 Skill</button></div>}

      <div className="setup-grid">
        <section className="panel resume-panel">
          <div className="section-heading"><span className="heading-icon mint"><FileText size={19} /></span><div><h2>先让我认识你</h2><p>简历可选，上传后会优先匹配你的技术经历</p></div><span className="optional">可选</span></div>
          <input ref={inputRef} hidden type="file" accept=".pdf,.docx,application/pdf,application/vnd.openxmlformats-officedocument.wordprocessingml.document" onChange={(event) => { const file = event.target.files?.[0]; event.target.value = ''; upload(file) }} />
          {!resume ? (
            <button className={`dropzone ${dragging ? 'dragging' : ''}`} onClick={() => inputRef.current?.click()} onDragOver={(event) => { event.preventDefault(); setDragging(true) }} onDragLeave={() => setDragging(false)} onDrop={(event) => { event.preventDefault(); setDragging(false); upload(event.dataTransfer.files[0]) }}>
              {uploading ? <LoaderCircle className="spin" size={28} /> : <UploadCloud size={28} />}
              <strong>{uploading ? '正在阅读简历…' : '拖入简历，或点击选择文件'}</strong><small>支持 PDF / DOCX · 最大 10 MB</small>
            </button>
          ) : (
            <div className="resume-result">
              <div className="file-row"><span className="file-icon"><FileCheck2 size={21} /></span><div><strong>{resume.fileName}</strong><small>已读取 {resume.characters.toLocaleString()} 个字符</small></div><button onClick={() => inputRef.current?.click()}>更换</button></div>
              <div className="keyword-title"><Sparkles size={14} />从简历中读到</div>
              <div className="keyword-cloud">{resume.keywords.length ? resume.keywords.map((keyword) => <span key={keyword}>{keyword}</span>) : <span className="muted-chip">暂未识别到技术关键词</span>}</div>
            </div>
          )}
        </section>

        <section className="panel direction-panel">
          <div className="section-heading"><span className="heading-icon violet"><Code2 size={19} /></span><div><h2>选择行业 / 领域 Skill</h2><p>每个 Skill 绑定独立知识库，后续可直接扩展新行业</p></div></div>
          <div className="industry-tabs">{(catalog?.industries ?? [{ id: 'computer', name: '计算机' }]).map((option) => <button key={option.id} className={industry === option.id ? 'selected' : ''} onClick={() => setIndustry(option.id)}>{option.name}</button>)}<span>更多行业 Skill 敬请期待</span></div>
          {catalogLoading ? <div className="skill-loading"><LoaderCircle className="spin" size={20} />正在加载领域 Skill…</div> : <div className="language-grid">{domainSkills.map((option) => (
            <button key={option.id} className={`language-card ${domainSkillId === option.id ? 'selected' : ''}`} onClick={() => setDomainSkillId(option.id)} style={{ '--language-color': option.accent } as CSSProperties}>
              <span className="language-badge">{option.shortLabel}</span><span><strong>{option.name}</strong><small>{option.topics?.slice(0, 3).join(' · ')}</small></span>{domainSkillId === option.id && <Check className="language-check" size={15} />}
            </button>
          ))}</div>}
        </section>

        <section className="panel interviewer-panel">
          <div className="section-heading"><span className="heading-icon mint"><Bot size={19} /></span><div><h2>选择面试官风格 Skill</h2><p>同一道题，不同面试官会用不同的评价重点和反馈语气</p></div></div>
          {catalogLoading ? <div className="skill-loading"><LoaderCircle className="spin" size={20} />正在加载面试官 Skill…</div> : <div className="interviewer-grid">{(catalog?.interviewers ?? []).map((skill) => <button key={skill.id} className={interviewerSkillId === skill.id ? 'selected' : ''} onClick={() => setInterviewerSkillId(skill.id)} style={{ '--skill-color': skill.accent } as CSSProperties}><span className="interviewer-avatar"><Bot size={18} /></span><strong>{skill.name}</strong><p>{skill.description}</p><small>{skill.evaluationFocus?.join(' · ')}</small>{interviewerSkillId === skill.id && <Check className="skill-check" size={15} />}</button>)}</div>}
        </section>

        <section className="panel preferences-panel">
          <div className="section-heading"><span className="heading-icon amber"><Gauge size={19} /></span><div><h2>定一下节奏</h2><p>不用紧张，你可以在回答前慢慢组织思路</p></div></div>
          <label className="field-label">怎么称呼你</label>
          <div className="name-input"><UserRound size={17} /><input value={candidateName} maxLength={30} onChange={(event) => setCandidateName(event.target.value)} placeholder="候选人（可不填）" /></div>
          <label className="field-label">面试难度</label>
          <div className="difficulty-row">{difficultyOptions.map((option) => <button key={option.value} className={difficulty === option.value ? 'selected' : ''} onClick={() => setDifficulty(option.value)}><strong>{option.label}</strong><small>{option.caption}</small></button>)}</div>
          <label className="foundation-toggle"><input type="checkbox" checked={includeFoundation} onChange={(event) => setIncludeFoundation(event.target.checked)} /><span><LibraryBig size={16} /><strong>混入计算机基础公共库</strong><small>每场加入 1–2 道操作系统、网络、数据库或分布式基础题</small></span><i /></label>
          <div className="count-row"><span><MessageSquareText size={16} /><span><strong>主问题数量</strong><small>每题还会连续追问 {followUpRounds} 轮</small></span></span><div className="count-stepper"><button onClick={() => setQuestionCount(Math.max(1, questionCount - 1))} disabled={questionCount <= 1}>−</button><input aria-label="主问题数量" type="number" min={1} max={maxQuestionCount} value={questionCount} onChange={(event) => setQuestionCount(Math.min(maxQuestionCount, Math.max(1, Number(event.target.value) || 1)))} /><button onClick={() => setQuestionCount(Math.min(maxQuestionCount, questionCount + 1))} disabled={questionCount >= maxQuestionCount}>＋</button><em>最多 {maxQuestionCount} 题</em></div></div>
          <button className="primary-action" onClick={start} disabled={starting || !catalog}>{starting ? <LoaderCircle className="spin" size={19} /> : <Sparkles size={18} />}开始模拟面试<ArrowRight size={18} /></button>
          <div className="start-hint"><span><Clock3 size={14} />约 {(questionCount * (followUpRounds + 1) + 1) * 2}–{(questionCount * (followUpRounds + 1) + 1) * 4} 分钟</span><span><Bot size={14} />1 次自我介绍 · {questionCount * (followUpRounds + 1)} 轮技术对话 · {modelReady ? 'AI 动态追问' : '本地智能追问'}</span></div>
        </section>

        <aside className="coach-card">
          <span className="coach-icon"><Lightbulb size={19} /></span><div><strong>一个小建议</strong><p>试着先讲结论，再讲原理，最后用项目经历收尾。清晰的结构往往比堆砌术语更有说服力。</p></div>
        </aside>
      </div>
    </div>
  )
}

function InterviewScreen({ session, turns, setTurns, setSession, onFinish, onBack, notify }: { session: Session; turns: Turn[]; setTurns: Dispatch<SetStateAction<Turn[]>>; setSession: (value: Session) => void; onFinish: (report: Report) => void; onBack: () => void; notify: (value: Toast) => void }) {
  const [answer, setAnswer] = useState('')
  const [submitting, setSubmitting] = useState(false)
  const [readyReport, setReadyReport] = useState<Report>()
  const [questionStartedAt, setQuestionStartedAt] = useState(Date.now())
  const [elapsed, setElapsed] = useState(0)
  const bottomRef = useRef<HTMLDivElement>(null)
  useEffect(() => { const timer = window.setInterval(() => setElapsed(Math.floor((Date.now() - new Date(session.startedAt).getTime()) / 1000)), 1000); return () => clearInterval(timer) }, [session.startedAt])
  useEffect(() => { bottomRef.current?.scrollIntoView({ behavior: 'smooth' }) }, [turns, session.currentQuestion, readyReport])

  const submit = async () => {
    const trimmed = answer.trim()
    if (!trimmed || submitting || !session.currentQuestion) return
    setSubmitting(true)
    try {
      const question = session.currentQuestion
      const result = await api.answer(session.id, trimmed, Math.floor((Date.now() - questionStartedAt) / 1000))
      setTurns((previous) => [...previous, { question, answer: trimmed, evaluation: result.evaluation, mainIndex: question.stage === 'introduction' ? 0 : session.current + 1 }])
      setAnswer('')
      if (result.completed && result.report) {
        setReadyReport(result.report)
        setSession({ ...session, current: result.current, followUpRound: 0, phase: 'completed', status: 'completed', currentQuestion: undefined })
      } else if (result.nextQuestion) {
        setSession({ ...session, current: result.current, phase: result.phase, followUpRound: result.followUpRound, followUpTotal: result.followUpTotal, currentQuestion: result.nextQuestion })
        setQuestionStartedAt(Date.now())
      }
    } catch (error) {
      notify({ type: 'error', message: error instanceof Error ? error.message : '回答提交失败，请重试' })
    } finally { setSubmitting(false) }
  }

  return (
    <div className="interview-page page-enter">
      <header className="interview-topbar">
        <button className="icon-button" onClick={onBack} title="退出本场面试"><ArrowLeft size={19} /></button>
        <div className="interview-title"><span className="live-dot" /><div><strong>{session.domainSkillName || languageLabel(session.language)}</strong><small>{session.candidateName} · {difficultyLabel(session.difficulty)} · {session.interviewerName}</small></div></div>
        <div className="progress-block"><div><span>{session.currentQuestion?.stage === 'introduction' ? '开场 · 自我介绍' : session.currentQuestion?.followUp ? `第 ${session.current + 1} 题 · 追问 ${session.currentQuestion.round}/${session.currentQuestion.followUpTotal}` : '主问题进度'}</span><strong>{session.currentQuestion?.stage === 'introduction' ? '热身' : `${Math.min(session.current + 1, session.total)} / ${session.total}`}</strong></div><div className="progress-track"><span style={{ width: `${session.currentQuestion?.stage === 'introduction' ? 3 : Math.min(100, ((session.current + (session.followUpRound / Math.max(1, session.followUpTotal + 1))) / session.total) * 100)}%` }} /></div></div>
        <div className="timer"><Clock3 size={16} />{formatTime(elapsed)}</div>
      </header>
      <div className="interview-layout">
        <section className="conversation">
          <div className="conversation-intro"><span><Bot size={21} /></span><div><strong>{session.interviewerName}</strong><p>你好，{session.candidateName}。{session.interviewerOpening}</p></div></div>
          {turns.map((turn) => <TurnCard key={turn.question.promptId || `${turn.question.id}-${turn.question.round}`} turn={turn} interviewerName={session.interviewerName} />)}
          {session.currentQuestion && (
            <div className="question-message">
              <div className="message-avatar"><Bot size={18} /></div>
              <div className={`message-content ${session.currentQuestion.followUp ? 'follow-up-question' : ''} ${session.currentQuestion.stage === 'introduction' ? 'introduction-question' : ''}`}><div className="message-meta"><strong>{session.interviewerName}</strong><span>{session.currentQuestion.stage === 'introduction' ? '开场 · 自我介绍' : session.currentQuestion.followUp ? `第 ${session.current + 1} 题 · 追问 ${session.currentQuestion.round}/${session.currentQuestion.followUpTotal}` : `第 ${session.current + 1} 题 · 主问题`}</span></div>{session.currentQuestion.leadIn && <p className="context-lead-in">{session.currentQuestion.leadIn}</p>}<p>{session.currentQuestion.prompt}</p><div className="question-tags">{session.currentQuestion.stage === 'introduction' && <span className="introduction-chip">不计分</span>}{session.currentQuestion.followUp && <span className="follow-up-chip">连续追问</span>}{session.currentQuestion.tags.map((tag) => <span key={tag}>{tag}</span>)}</div></div>
            </div>
          )}
          {readyReport && <div className="completion-card"><span><Sparkles size={26} /></span><h2>这场面试完成了</h2><p>你认真回答了 {readyReport.answered} 道题。标准答案、逐题评分和下一步练习建议都已经整理好。</p><button className="primary-action compact" onClick={() => onFinish(readyReport)}>查看我的复盘<ArrowRight size={18} /></button></div>}
          <div ref={bottomRef} />
        </section>
        <aside className="interview-aside">
          <div className="aside-card"><span className="aside-label">回答结构</span><div className="answer-framework"><div><span>1</span><p><strong>先讲结论</strong><small>一句话回应核心问题</small></p></div><div><span>2</span><p><strong>拆解原理</strong><small>说清机制与边界</small></p></div><div><span>3</span><p><strong>联系实践</strong><small>用项目或反例收尾</small></p></div></div></div>
          <div className="aside-card quiet"><Target size={18} /><p>不确定时可以明确假设，再沿着假设推理。面试官也在观察你的思考过程。</p></div>
        </aside>
      </div>
      {session.currentQuestion && (
        <div className="composer-wrap"><div className="composer"><textarea autoFocus value={answer} maxLength={8000} onChange={(event) => setAnswer(event.target.value)} onKeyDown={(event) => { if ((event.metaKey || event.ctrlKey) && event.key === 'Enter') submit() }} placeholder={session.currentQuestion.stage === 'introduction' ? '用 1–2 分钟介绍自己… 可以从技术方向、项目职责和成果讲起' : session.currentQuestion.followUp ? `回应第 ${session.currentQuestion.round} 轮追问… 延续刚才的思路，补充边界或案例` : '说说你的思路… 可以先写结论，再补充细节'} /><div className="composer-footer"><span>{answer.length > 0 ? `${answer.length} 字` : '⌘ / Ctrl + Enter 发送'}</span><button onClick={submit} disabled={!answer.trim() || submitting}>{submitting ? <LoaderCircle className="spin" size={18} /> : <Send size={17} />}{submitting ? (session.currentQuestion.stage === 'introduction' ? '正在结合简历准备问题' : '正在点评并准备追问') : (session.currentQuestion.stage === 'introduction' ? '完成自我介绍' : '提交回答')}</button></div></div></div>
      )}
    </div>
  )
}

function TurnCard({ turn, interviewerName }: { turn: Turn; interviewerName: string }) {
  const [open, setOpen] = useState(false)
  const introduction = turn.question.stage === 'introduction'
  return (
    <div className="turn-block">
      <div className="question-message completed-question"><div className="message-avatar"><Bot size={18} /></div><div className={`message-content ${turn.question.followUp ? 'follow-up-question' : ''} ${introduction ? 'introduction-question' : ''}`}><div className="message-meta"><strong>{interviewerName}</strong><span>{introduction ? '开场 · 自我介绍' : turn.question.followUp ? `第 ${turn.mainIndex} 题 · 追问 ${turn.question.round}/${turn.question.followUpTotal}` : `第 ${turn.mainIndex} 题 · 主问题`}</span></div><p>{turn.question.prompt}</p></div></div>
      <div className="answer-message"><div className="answer-bubble"><p>{turn.answer}</p></div><div className="user-avatar"><UserRound size={17} /></div></div>
      {introduction ? <div className="introduction-recorded"><Check size={15} /><span><strong>自我介绍已记录</strong><small>不计入技术得分，后续问题会结合你的经历展开</small></span></div> : <button className={`inline-evaluation ${open ? 'open' : ''}`} onClick={() => setOpen(!open)}><span className={`score-dot score-${scoreBand(turn.evaluation.score)}`}>{turn.evaluation.score}</span><span><strong>{turn.evaluation.summary}</strong><small>{turn.evaluation.source === 'llm' ? 'Eino AI 点评' : '本地关键点评分'} · 点击{open ? '收起' : '展开'}</small></span><ChevronDown size={18} />
        {open && <div className="evaluation-details"><div><Check size={15} /><p>{turn.evaluation.strengths.join('；')}</p></div><div><Zap size={15} /><p>{turn.evaluation.improvements.join('；')}</p></div></div>}
      </button>}
    </div>
  )
}

function ReviewScreen({ report, onRestart }: { report: Report; onRestart: () => void }) {
  const [openIndex, setOpenIndex] = useState(0)
  return (
    <div className="review-page page-enter">
      <header className="review-hero"><div><span className="eyebrow"><FileCheck2 size={14} /> 本场面试复盘</span><h1>{report.candidateName}，你已经走完了这一轮。</h1><p>{report.domainSkillName || languageLabel(report.language)} · {report.interviewerName} · {difficultyLabel(report.difficulty)} · {report.answered} 道题 · {formatTime(report.durationSeconds)}</p></div><button className="secondary-action" onClick={onRestart}><RefreshCw size={17} />再练一场</button></header>
      <div className="score-overview">
        <div className="score-ring" style={{ '--score': report.score } as CSSProperties}><div><strong>{report.score}</strong><span>综合得分</span></div></div>
        <div className="score-copy"><span className={`performance-pill ${scoreBand(report.score)}`}>{performanceText(report.score)}</span><h2>{report.score >= 80 ? '表达清晰，继续保持工程细节。' : report.score >= 60 ? '基本盘不错，再补齐关键边界。' : '已经找到薄弱点，这正是练习的价值。'}</h2><p>得分由逐题关键点或 Eino 模型评价汇总。建议先看“值得保留”，再逐题对照标准答案。</p></div>
        <div className="summary-columns"><div><span className="summary-icon good"><Check size={17} /></span><section><strong>值得保留</strong>{report.highlights.length ? report.highlights.map((item) => <p key={item}>{item}</p>) : <p>完成了全部作答</p>}</section></div><div><span className="summary-icon focus"><Target size={17} /></span><section><strong>下轮重点</strong>{report.focusAreas.length ? report.focusAreas.map((item) => <p key={item}>{item}</p>) : <p>多用项目案例支撑结论</p>}</section></div></div>
      </div>
      {report.introduction && <section className="introduction-review"><div><span><UserRound size={17} /></span><div><strong>本场自我介绍</strong><small>开场记录 · 不计入技术得分</small></div></div><p>{report.introduction.answer}</p></section>}
      <section className="review-list"><div className="review-heading"><div><h2>逐题回看</h2><p>主回答、连续追问、点评与标准答案都在这里</p></div><span>1 次自我介绍 · {report.answers.length} 个主问题 · {report.answers.reduce((total, record) => total + 1 + (record.followUps?.length ?? 0), 0)} 轮技术对话</span></div>
        {report.answers.map((record, index) => {
          const score = record.averageScore || record.evaluation.score
          return <article key={record.question.id} className={`review-item ${openIndex === index ? 'open' : ''}`}>
            <button className="review-item-head" onClick={() => setOpenIndex(openIndex === index ? -1 : index)}><span className="question-number">{String(index + 1).padStart(2, '0')}</span><div><h3>{record.question.prompt}</h3><span>{record.question.tags.join(' · ')} · {(record.followUps?.length ?? 0)} 轮追问</span></div><span className={`review-score ${scoreBand(score)}`}>{score} 分</span><ChevronRight size={19} /></button>
            {openIndex === index && <div className="review-body"><div className="answer-compare"><section><span className="compare-label"><UserRound size={15} />主回答</span><p>{record.answer}</p></section><section className="standard"><span className="compare-label"><Sparkles size={15} />标准答案</span><p>{record.question.standardAnswer}</p><div className="key-points">{record.question.keyPoints.map((point) => <span key={point}>{point.split('/')[0]}</span>)}</div></section></div><div className="coach-feedback"><BrainCircuit size={18} /><div><strong>主回答点评</strong><p>{record.evaluation.summary} {record.evaluation.improvements.join('；')}</p></div></div>{record.followUps?.length > 0 && <div className="follow-up-review"><span className="follow-up-review-title"><MessageSquareText size={15} />连续追问记录</span>{record.followUps.map((followUp) => <section key={followUp.round}><div><strong>追问 {followUp.round}</strong><em>{followUp.evaluation.score} 分</em></div><h4>{followUp.prompt}</h4><p><span>你的回答</span>{followUp.answer}</p><small>{followUp.evaluation.summary} {followUp.evaluation.improvements.join('；')}</small></section>)}</div>}</div>}
          </article>
        })}
      </section>
      <div className="review-footer"><button className="primary-action compact" onClick={onRestart}><RefreshCw size={17} />开始下一场</button></div>
    </div>
  )
}

function LearningPage({ catalog, notify }: { catalog?: SkillCatalog; notify: (value: Toast) => void }) {
  const [resources, setResources] = useState<LearningResource[]>([])
  const [domain, setDomain] = useState('all')
  const [kind, setKind] = useState('all')
  const [selectedOnly, setSelectedOnly] = useState(false)
  const [loading, setLoading] = useState(true)
  const [refreshing, setRefreshing] = useState(false)
  const [refreshedAt, setRefreshedAt] = useState<string>()
  const load = async () => {
    setLoading(true)
    try { const result = await api.learningResources(domain, kind, selectedOnly); setResources(result.resources); setRefreshedAt(result.refreshedAt) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '学习资源加载失败' }) }
    finally { setLoading(false) }
  }
  useEffect(() => { load() }, [domain, kind, selectedOnly])
  const refresh = async () => {
    setRefreshing(true)
    try {
      const result = await api.refreshLearning()
      setRefreshedAt(result.refreshedAt)
      if (result.warnings?.length) notify({ type: 'error', message: `部分来源暂不可用，其余已更新（${result.warnings.length} 个）` })
      else notify({ type: 'success', message: '已从权威来源抓取最新文章' })
      await load()
    } catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '抓取失败' }) }
    finally { setRefreshing(false) }
  }
  const toggle = async (resource: LearningResource) => {
    const selected = !resource.selected
    setResources((previous) => previous.map((item) => item.id === resource.id ? { ...item, selected } : item))
    try { await api.selectLearning(resource.id, selected); notify({ type: 'success', message: selected ? '已加入我的学习清单' : '已移出学习清单' }) }
    catch (error) { setResources((previous) => previous.map((item) => item.id === resource.id ? resource : item)); notify({ type: 'error', message: error instanceof Error ? error.message : '操作失败' }) }
  }
  const domainSkills = catalog?.domains ?? []
  return <div className="learning-page page-enter">
    <header className="library-hero"><div><span className="eyebrow"><BookOpen size={14} /> Learning Studio</span><h1>从权威内容里，<em>继续生长。</em></h1><p>官方文档、技术博客、课程和视频集中在这里。只抓取标题与摘要，阅读始终回到原始权威网站。</p></div><button className="secondary-action" onClick={refresh} disabled={refreshing}>{refreshing ? <LoaderCircle className="spin" size={17} /> : <RefreshCw size={17} />}抓取最新内容</button></header>
    <div className="learning-toolbar"><div className="filter-group"><button className={domain === 'all' ? 'selected' : ''} onClick={() => setDomain('all')}>全部领域</button>{domainSkills.map((skill) => <button key={skill.id} className={domain === skill.id ? 'selected' : ''} onClick={() => setDomain(skill.id)}>{skill.shortLabel}</button>)}</div><div className="filter-group kind-filter">{['all', 'article', 'video', 'course', 'docs'].map((value) => <button key={value} className={kind === value ? 'selected' : ''} onClick={() => setKind(value)}>{resourceKindLabel(value)}</button>)}</div><label className="selected-filter"><input type="checkbox" checked={selectedOnly} onChange={(event) => setSelectedOnly(event.target.checked)} /><BookmarkCheck size={15} />只看学习清单</label></div>
    <div className="resource-meta"><span><Globe2 size={14} />{resources.length} 个权威资源</span>{refreshedAt && <span>上次抓取 {new Date(refreshedAt).toLocaleString('zh-CN')}</span>}</div>
    {loading ? <div className="empty-state"><LoaderCircle className="spin" size={26} /><p>正在整理学习内容…</p></div> : resources.length === 0 ? <div className="empty-state"><BookOpen size={28} /><h3>这里还没有内容</h3><p>调整筛选条件，或点击“抓取最新内容”。</p></div> : <div className="resource-grid">{resources.map((resource) => <article key={resource.id} className={`resource-card ${resource.selected ? 'selected' : ''}`}><div className="resource-card-top"><span className={`resource-kind kind-${resource.kind}`}>{resourceKindIcon(resource.kind)}{resourceKindLabel(resource.kind)}</span>{resource.live && <span className="live-source"><span />实时抓取</span>}</div><h2>{resource.title}</h2><p>{resource.summary || '来自权威来源的最新内容，点击前往原站学习。'}</p><div className="resource-source"><strong>{resource.authority}</strong><small>{resourceDate(resource)}</small></div><footer><button className={`bookmark-button ${resource.selected ? 'active' : ''}`} onClick={() => toggle(resource)}>{resource.selected ? <BookmarkCheck size={15} /> : <Bookmark size={15} />}{resource.selected ? '已选择' : '加入学习'}</button><a href={resource.url} target="_blank" rel="noreferrer">去学习<ArrowRight size={15} /></a></footer></article>)}</div>}
  </div>
}

function KnowledgePage({ notify }: { notify: (value: Toast) => void }) {
  const [bases, setBases] = useState<KnowledgeBase[]>([])
  const [baseId, setBaseId] = useState('')
  const [items, setItems] = useState<QAItem[]>([])
  const [query, setQuery] = useState('')
  const [loading, setLoading] = useState(true)
  const [openItem, setOpenItem] = useState<string>()
  const [adding, setAdding] = useState(false)
  const [question, setQuestion] = useState('')
  const [answer, setAnswer] = useState('')
  const loadBases = async () => {
    try { const result = await api.knowledgeBases(); setBases(result.bases); if (!baseId && result.bases.length) setBaseId(result.bases[0].id) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '知识库加载失败' }) }
  }
  const search = async (targetBase = baseId, targetQuery = query) => {
    if (!targetBase) return
    setLoading(true)
    try { const result = await api.searchKnowledge(targetBase, targetQuery); setItems(result.items); setOpenItem(undefined) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '检索失败' }) }
    finally { setLoading(false) }
  }
  useEffect(() => { loadBases() }, [])
  useEffect(() => { if (baseId) search(baseId, '') }, [baseId])
  const addQA = async () => {
    try { await api.addKnowledge(baseId, { question, answer, difficulty: 'medium' }); setQuestion(''); setAnswer(''); setAdding(false); await Promise.all([search(), loadBases()]); notify({ type: 'success', message: 'QA 已写入当前知识库' }) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '写入失败' }) }
  }
  const currentBase = bases.find((base) => base.id === baseId)
  return <div className="knowledge-page page-enter">
    <header className="library-hero"><div><span className="eyebrow"><LibraryBig size={14} /> Knowledge Base</span><h1>每次出题，都让知识库<em>更好检索。</em></h1><p>语言库彼此独立，计算机基础作为公共库。每次面试实际出过的题都会累计出题次数并沉淀为 QA。</p></div><button className="secondary-action" onClick={() => setAdding(!adding)}>{adding ? <X size={17} /> : <Sparkles size={17} />}{adding ? '取消' : '新增 QA'}</button></header>
    <div className="knowledge-layout"><aside className="base-list"><span className="aside-label">知识库</span>{bases.map((base) => <button key={base.id} className={baseId === base.id ? 'selected' : ''} onClick={() => { setBaseId(base.id); setQuery('') }} style={{ '--base-color': base.accent } as CSSProperties}><span>{base.language === 'foundation' ? 'CS' : languageOptions.find((option) => option.value === base.language)?.short ?? base.language}</span><div><strong>{base.name}</strong><small>{base.itemCount} 个 QA · 已出题 {base.issuedCount} 次</small></div></button>)}</aside><section className="knowledge-main">
      {currentBase && <div className="base-summary"><span className="heading-icon violet"><LibraryBig size={19} /></span><div><h2>{currentBase.name}</h2><p>{currentBase.description}</p><div>{currentBase.topics.map((topic) => <span key={topic}>{topic}</span>)}</div></div></div>}
      {adding && <div className="qa-form"><label>问题<input value={question} onChange={(event) => setQuestion(event.target.value)} placeholder="输入希望沉淀的问题" /></label><label>标准答案<textarea value={answer} onChange={(event) => setAnswer(event.target.value)} placeholder="输入可复盘、可检索的标准答案" /></label><button className="primary-action compact" disabled={!question.trim() || !answer.trim()} onClick={addQA}><Check size={16} />写入知识库</button></div>}
      <div className="knowledge-search"><div><SearchIcon /><input value={query} onChange={(event) => setQuery(event.target.value)} onKeyDown={(event) => { if (event.key === 'Enter') search() }} placeholder="检索题目、答案或标签…" /></div><button onClick={() => search()}>检索</button></div>
      {loading ? <div className="empty-state compact"><LoaderCircle className="spin" size={23} /></div> : <div className="qa-list">{items.map((item) => <article key={item.id} className={openItem === item.id ? 'open' : ''}><button onClick={() => setOpenItem(openItem === item.id ? undefined : item.id)}><span className="qa-source">{item.source === 'manual' ? '手动' : item.source === 'interview' ? '面试沉淀' : '内置'}</span><div><h3>{item.question}</h3><small>{item.tags.join(' · ')}{item.issuedCount > 0 && ` · 已出题 ${item.issuedCount} 次`}</small></div><ChevronDown size={18} /></button>{openItem === item.id && <div className="qa-answer"><span>标准答案</span><p>{item.answer}</p><div>{item.keyPoints.map((point) => <em key={point}>{point.split('/')[0]}</em>)}</div></div>}</article>)}</div>}
    </section></div>
  </div>
}

function SearchIcon() { return <Target size={16} /> }
function resourceKindLabel(kind: string) { return ({ all: '全部类型', article: '文章', video: '视频', course: '课程', docs: '文档' } as Record<string, string>)[kind] ?? kind }
function resourceKindIcon(kind: string) { if (kind === 'video') return <PlayCircle size={13} />; if (kind === 'article') return <Newspaper size={13} />; return <BookOpen size={13} /> }
function resourceDate(resource: LearningResource) { const date = resource.publishedAt ? new Date(resource.publishedAt) : undefined; return date && date.getFullYear() > 1900 ? date.toLocaleDateString('zh-CN') : resource.source }

function SettingsDrawer({ open, value, onClose, onSaved, notify }: { open: boolean; value: ModelConfig; onClose: () => void; onSaved: (value: ModelConfig) => void; notify: (value: Toast) => void }) {
  const [apiKey, setApiKey] = useState('')
  const [baseUrl, setBaseUrl] = useState(value.baseUrl)
  const [model, setModel] = useState(value.model)
  const [enabled, setEnabled] = useState(value.enabled)
  const [saving, setSaving] = useState(false)
  const [testing, setTesting] = useState(false)
  useEffect(() => { if (open) { setApiKey(''); setBaseUrl(value.baseUrl); setModel(value.model); setEnabled(value.enabled) } }, [open, value])
  if (!open) return null
  const save = async () => {
    setSaving(true)
    try { const next = await api.saveModelConfig({ apiKey, baseUrl, model, enabled }); onSaved(next); notify({ type: 'success', message: '模型配置已保存' }); onClose() }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '保存失败' }) }
    finally { setSaving(false) }
  }
  const test = async () => {
    setTesting(true)
    try {
      const next = await api.saveModelConfig({ apiKey, baseUrl, model, enabled: true }); onSaved(next)
      const result = await api.testModelConfig(); setEnabled(true); notify({ type: 'success', message: result.message })
    } catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '连接失败' }) }
    finally { setTesting(false) }
  }
  return <div className="drawer-backdrop" onMouseDown={(event) => { if (event.currentTarget === event.target) onClose() }}><aside className="settings-drawer"><header><div><span className="heading-icon violet"><Settings2 size={19} /></span><section><h2>大模型设置</h2><p>通过 Eino 接入 OpenAI-compatible API</p></section></div><button className="icon-button" onClick={onClose}><X size={19} /></button></header><div className="drawer-content"><div className="settings-callout"><ShieldCheck size={19} /><p><strong>凭据保存在本机</strong><br />API Key 以 0600 权限写入应用数据目录，不会进入项目代码或日志。</p></div><label className="setting-field"><span>API Key {value.hasApiKey && <em>已保存</em>}</span><input type="password" autoComplete="off" value={apiKey} onChange={(event) => setApiKey(event.target.value)} placeholder={value.hasApiKey ? '留空以继续使用已保存的 Key' : 'sk-...'} /></label><label className="setting-field"><span>Base URL <small>可留空使用默认地址</small></span><input value={baseUrl} onChange={(event) => setBaseUrl(event.target.value)} placeholder="https://api.example.com/v1" /></label><label className="setting-field"><span>Model</span><input value={model} onChange={(event) => setModel(event.target.value)} placeholder="模型 ID" /></label><label className="toggle-row"><span><strong>启用 AI 深度点评</strong><small>不可用时会自动降级为本地评分</small></span><input type="checkbox" checked={enabled} onChange={(event) => setEnabled(event.target.checked)} /><i /></label><div className="settings-note"><Bot size={16} /><p>模型只接收当前题目、标准答案和你的当前回答，用于生成更细致的评分建议。</p></div></div><footer><button className="secondary-action" onClick={test} disabled={testing || (!apiKey && !value.hasApiKey) || !model}>{testing ? <LoaderCircle className="spin" size={16} /> : <Zap size={16} />}测试连接</button><button className="primary-action compact" onClick={save} disabled={saving}>{saving ? <LoaderCircle className="spin" size={16} /> : <Check size={16} />}保存设置</button></footer></aside></div>
}

function languageLabel(language: Language) { return languageOptions.find((option) => option.value === language)?.label ?? language }
function difficultyLabel(value: Difficulty) { return difficultyOptions.find((option) => option.value === value)?.label ?? value }
function formatTime(seconds: number) { const minutes = Math.floor(seconds / 60); return `${String(minutes).padStart(2, '0')}:${String(seconds % 60).padStart(2, '0')}` }
function scoreBand(score: number) { return score >= 80 ? 'great' : score >= 60 ? 'good' : 'focus' }
function performanceText(score: number) { return score >= 80 ? '表现出色' : score >= 60 ? '稳步进阶' : '继续加油' }

export default App
