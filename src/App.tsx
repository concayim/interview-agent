import { useEffect, useRef, useState } from 'react'
import type { CSSProperties, Dispatch, SetStateAction } from 'react'
import {
  ArrowLeft,
  ArrowRight,
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
  Gauge,
  Lightbulb,
  LoaderCircle,
  MessageSquareText,
  RefreshCw,
  Send,
  Settings2,
  ShieldCheck,
  Sparkles,
  Target,
  UploadCloud,
  UserRound,
  X,
  Zap,
} from 'lucide-react'
import { api } from './api'
import type { Difficulty, Evaluation, Language, ModelConfig, Question, Report, Resume, Session } from './types'

type Screen = 'setup' | 'interview' | 'review'
type Turn = { question: Question; answer: string; evaluation: Evaluation }
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
  const [toast, setToast] = useState<Toast>()

  useEffect(() => {
    api.getModelConfig().then(setModelConfig).catch((error) => setToast({ type: 'error', message: error.message }))
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
      <Sidebar screen={screen} modelReady={modelConfig.enabled && modelConfig.hasApiKey} onHome={restart} onSettings={() => setSettingsOpen(true)} />
      <main className="app-main">
        {screen === 'setup' && <Setup resume={resume} onResume={setResume} onStart={start} modelReady={modelConfig.enabled && modelConfig.hasApiKey} notify={setToast} />}
        {screen === 'interview' && session && (
          <InterviewScreen session={session} turns={turns} setTurns={setTurns} setSession={setSession} onFinish={finish} onBack={restart} notify={setToast} />
        )}
        {screen === 'review' && report && <ReviewScreen report={report} onRestart={restart} />}
      </main>
      <SettingsDrawer open={settingsOpen} value={modelConfig} onClose={() => setSettingsOpen(false)} onSaved={setModelConfig} notify={setToast} />
      {toast && <div className={`toast toast-${toast.type}`}>{toast.type === 'success' ? <Check size={18} /> : <CircleAlert size={18} />}{toast.message}</div>}
    </div>
  )
}

function Sidebar({ screen, modelReady, onHome, onSettings }: { screen: Screen; modelReady: boolean; onHome: () => void; onSettings: () => void }) {
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
      <div className={`model-pill ${modelReady ? 'online' : ''}`}><span className="status-dot" /><div><strong>{modelReady ? 'AI 增强已开启' : '本地模式'}</strong><small>{modelReady ? 'Eino 模型在线' : '规则评分可用'}</small></div></div>
      <button className="sidebar-button" onClick={onSettings}><Settings2 size={17} />模型设置</button>
      <div className="privacy-note"><ShieldCheck size={15} /><span>简历仅在本机解析</span></div>
    </aside>
  )
}

function Setup({ resume, onResume, onStart, modelReady, notify }: { resume?: Resume; onResume: (value: Resume) => void; onStart: (value: Session) => void; modelReady: boolean; notify: (value: Toast) => void }) {
  const [candidateName, setCandidateName] = useState('')
  const [language, setLanguage] = useState<Language>('golang')
  const [difficulty, setDifficulty] = useState<Difficulty>('mixed')
  const [questionCount, setQuestionCount] = useState(5)
  const [uploading, setUploading] = useState(false)
  const [starting, setStarting] = useState(false)
  const [dragging, setDragging] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (difficulty !== 'mixed') setQuestionCount(2)
    else if (questionCount === 2) setQuestionCount(5)
  }, [difficulty])

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
    try { onStart(await api.startInterview({ candidateName, resumeId: resume?.id, language, difficulty, questionCount })) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '面试创建失败' }) }
    finally { setStarting(false) }
  }

  return (
    <div className="setup-page page-enter">
      <header className="hero">
        <div><span className="eyebrow"><Sparkles size={14} /> AI 技术面试陪练</span><h1>把每一次回答，<br /><em>练成底气。</em></h1><p>上传简历，选择技术方向。面试官会围绕你的经历发问，并在最后给出有标准答案的完整复盘。</p></div>
        <div className="hero-orbit" aria-hidden="true"><span className="orbit orbit-one" /><span className="orbit orbit-two" /><BrainCircuit size={56} /></div>
      </header>

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
          <div className="section-heading"><span className="heading-icon violet"><Code2 size={19} /></span><div><h2>选择技术方向</h2><p>每套题都有人工编写的参考答案与关键点</p></div></div>
          <div className="language-grid">{languageOptions.map((option) => (
            <button key={option.value} className={`language-card ${language === option.value ? 'selected' : ''}`} onClick={() => setLanguage(option.value)} style={{ '--language-color': option.color } as CSSProperties}>
              <span className="language-badge">{option.short}</span><span><strong>{option.label}</strong><small>{option.caption}</small></span>{language === option.value && <Check className="language-check" size={15} />}
            </button>
          ))}</div>
        </section>

        <section className="panel preferences-panel">
          <div className="section-heading"><span className="heading-icon amber"><Gauge size={19} /></span><div><h2>定一下节奏</h2><p>不用紧张，你可以在回答前慢慢组织思路</p></div></div>
          <label className="field-label">怎么称呼你</label>
          <div className="name-input"><UserRound size={17} /><input value={candidateName} maxLength={30} onChange={(event) => setCandidateName(event.target.value)} placeholder="候选人（可不填）" /></div>
          <label className="field-label">面试难度</label>
          <div className="difficulty-row">{difficultyOptions.map((option) => <button key={option.value} className={difficulty === option.value ? 'selected' : ''} onClick={() => setDifficulty(option.value)}><strong>{option.label}</strong><small>{option.caption}</small></button>)}</div>
          {difficulty === 'mixed' && <div className="count-row"><span><MessageSquareText size={16} />题目数量</span><div>{[3, 5, 6].map((count) => <button key={count} className={questionCount === count ? 'selected' : ''} onClick={() => setQuestionCount(count)}>{count} 题</button>)}</div></div>}
          <button className="primary-action" onClick={start} disabled={starting}>{starting ? <LoaderCircle className="spin" size={19} /> : <Sparkles size={18} />}开始模拟面试<ArrowRight size={18} /></button>
          <div className="start-hint"><span><Clock3 size={14} />约 {questionCount * 3}–{questionCount * 5} 分钟</span><span><Bot size={14} />{modelReady ? 'AI 深度点评' : '本地关键点评分'}</span></div>
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
      setTurns((previous) => [...previous, { question, answer: trimmed, evaluation: result.evaluation }])
      setAnswer('')
      if (result.completed && result.report) {
        setReadyReport(result.report)
        setSession({ ...session, current: result.current, status: 'completed', currentQuestion: undefined })
      } else if (result.nextQuestion) {
        setSession({ ...session, current: result.current, currentQuestion: result.nextQuestion })
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
        <div className="interview-title"><span className="live-dot" /><div><strong>{languageLabel(session.language)} 技术面试</strong><small>{session.candidateName} · {difficultyLabel(session.difficulty)}</small></div></div>
        <div className="progress-block"><div><span>进度</span><strong>{Math.min(session.current + 1, session.total)} / {session.total}</strong></div><div className="progress-track"><span style={{ width: `${Math.min(100, (session.current / session.total) * 100)}%` }} /></div></div>
        <div className="timer"><Clock3 size={16} />{formatTime(elapsed)}</div>
      </header>
      <div className="interview-layout">
        <section className="conversation">
          <div className="conversation-intro"><span><Bot size={21} /></span><div><strong>面试官 Echo</strong><p>你好，{session.candidateName}。我们从 {languageLabel(session.language)} 开始。别急着给“标准句式”，我更想听你的理解和取舍。</p></div></div>
          {turns.map((turn, index) => <TurnCard key={turn.question.id} turn={turn} index={index + 1} />)}
          {session.currentQuestion && (
            <div className="question-message">
              <div className="message-avatar"><Bot size={18} /></div>
              <div className="message-content"><div className="message-meta"><strong>Echo</strong><span>第 {session.current + 1} 题</span></div><p>{session.currentQuestion.prompt}</p><div className="question-tags">{session.currentQuestion.tags.map((tag) => <span key={tag}>{tag}</span>)}</div></div>
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
        <div className="composer-wrap"><div className="composer"><textarea autoFocus value={answer} maxLength={8000} onChange={(event) => setAnswer(event.target.value)} onKeyDown={(event) => { if ((event.metaKey || event.ctrlKey) && event.key === 'Enter') submit() }} placeholder="说说你的思路… 可以先写结论，再补充细节" /><div className="composer-footer"><span>{answer.length > 0 ? `${answer.length} 字` : '⌘ / Ctrl + Enter 发送'}</span><button onClick={submit} disabled={!answer.trim() || submitting}>{submitting ? <LoaderCircle className="spin" size={18} /> : <Send size={17} />}{submitting ? '正在点评' : '提交回答'}</button></div></div></div>
      )}
    </div>
  )
}

function TurnCard({ turn, index }: { turn: Turn; index: number }) {
  const [open, setOpen] = useState(false)
  return (
    <div className="turn-block">
      <div className="question-message completed-question"><div className="message-avatar"><Bot size={18} /></div><div className="message-content"><div className="message-meta"><strong>Echo</strong><span>第 {index} 题</span></div><p>{turn.question.prompt}</p></div></div>
      <div className="answer-message"><div className="answer-bubble"><p>{turn.answer}</p></div><div className="user-avatar"><UserRound size={17} /></div></div>
      <button className={`inline-evaluation ${open ? 'open' : ''}`} onClick={() => setOpen(!open)}><span className={`score-dot score-${scoreBand(turn.evaluation.score)}`}>{turn.evaluation.score}</span><span><strong>{turn.evaluation.summary}</strong><small>{turn.evaluation.source === 'llm' ? 'Eino AI 点评' : '本地关键点评分'} · 点击{open ? '收起' : '展开'}</small></span><ChevronDown size={18} />
        {open && <div className="evaluation-details"><div><Check size={15} /><p>{turn.evaluation.strengths.join('；')}</p></div><div><Zap size={15} /><p>{turn.evaluation.improvements.join('；')}</p></div></div>}
      </button>
    </div>
  )
}

function ReviewScreen({ report, onRestart }: { report: Report; onRestart: () => void }) {
  const [openIndex, setOpenIndex] = useState(0)
  return (
    <div className="review-page page-enter">
      <header className="review-hero"><div><span className="eyebrow"><FileCheck2 size={14} /> 本场面试复盘</span><h1>{report.candidateName}，你已经走完了这一轮。</h1><p>{languageLabel(report.language)} · {difficultyLabel(report.difficulty)} · {report.answered} 道题 · {formatTime(report.durationSeconds)}</p></div><button className="secondary-action" onClick={onRestart}><RefreshCw size={17} />再练一场</button></header>
      <div className="score-overview">
        <div className="score-ring" style={{ '--score': report.score } as CSSProperties}><div><strong>{report.score}</strong><span>综合得分</span></div></div>
        <div className="score-copy"><span className={`performance-pill ${scoreBand(report.score)}`}>{performanceText(report.score)}</span><h2>{report.score >= 80 ? '表达清晰，继续保持工程细节。' : report.score >= 60 ? '基本盘不错，再补齐关键边界。' : '已经找到薄弱点，这正是练习的价值。'}</h2><p>得分由逐题关键点或 Eino 模型评价汇总。建议先看“值得保留”，再逐题对照标准答案。</p></div>
        <div className="summary-columns"><div><span className="summary-icon good"><Check size={17} /></span><section><strong>值得保留</strong>{report.highlights.length ? report.highlights.map((item) => <p key={item}>{item}</p>) : <p>完成了全部作答</p>}</section></div><div><span className="summary-icon focus"><Target size={17} /></span><section><strong>下轮重点</strong>{report.focusAreas.length ? report.focusAreas.map((item) => <p key={item}>{item}</p>) : <p>多用项目案例支撑结论</p>}</section></div></div>
      </div>
      <section className="review-list"><div className="review-heading"><div><h2>逐题回看</h2><p>你的回答、点评与标准答案都在这里</p></div><span>{report.answers.length} 个回答</span></div>
        {report.answers.map((record, index) => (
          <article key={record.question.id} className={`review-item ${openIndex === index ? 'open' : ''}`}>
            <button className="review-item-head" onClick={() => setOpenIndex(openIndex === index ? -1 : index)}><span className="question-number">{String(index + 1).padStart(2, '0')}</span><div><h3>{record.question.prompt}</h3><span>{record.question.tags.join(' · ')}</span></div><span className={`review-score ${scoreBand(record.evaluation.score)}`}>{record.evaluation.score} 分</span><ChevronRight size={19} /></button>
            {openIndex === index && <div className="review-body"><div className="answer-compare"><section><span className="compare-label"><UserRound size={15} />你的回答</span><p>{record.answer}</p></section><section className="standard"><span className="compare-label"><Sparkles size={15} />标准答案</span><p>{record.question.standardAnswer}</p><div className="key-points">{record.question.keyPoints.map((point) => <span key={point}>{point.split('/')[0]}</span>)}</div></section></div><div className="coach-feedback"><BrainCircuit size={18} /><div><strong>面试官点评</strong><p>{record.evaluation.summary} {record.evaluation.improvements.join('；')}</p></div></div></div>}
          </article>
        ))}
      </section>
      <div className="review-footer"><button className="primary-action compact" onClick={onRestart}><RefreshCw size={17} />开始下一场</button></div>
    </div>
  )
}

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
