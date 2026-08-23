import { lazy, Suspense, useEffect, useRef, useState } from 'react'
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
  Mic,
  Newspaper,
  PlayCircle,
  RefreshCw,
  Search,
  Send,
  Settings2,
  ShieldCheck,
  Sparkles,
  Target,
  LibraryBig,
  UploadCloud,
  UserRound,
  Video,
  VideoOff,
  Volume2,
  VolumeX,
  X,
  Zap,
} from 'lucide-react'
import { api } from './api'
import { startRealtimeSpeech } from './speech'
import type { RealtimeSpeechSession } from './speech'
import { InterviewVoicePlayer } from './tts'
import { deriveAvatarMode } from './avatar-state'
import { buildCompletionSpeech, buildFollowUpSpeech, buildInterviewOpeningSpeech, buildNextQuestionSpeech, isEnglishInterview } from './interview-language'
import type { Difficulty, Evaluation, KnowledgeBase, Language, LearningResource, ModelConfig, QAItem, Question, Report, Resume, Session, Skill, SkillCatalog } from './types'

type Screen = 'setup' | 'interview' | 'review' | 'learning' | 'knowledge'
type IntentMessage = { id: string; role: 'candidate' | 'assistant'; text: string; intent: string }
type Turn = { question: Question; questionNumber: number; followUp: boolean; answer: string; evaluation: Evaluation; sideMessages?: IntentMessage[] }
type Toast = { type: 'success' | 'error'; message: string }
type SpeechMode = 'idle' | 'connecting' | 'recording' | 'finalizing'

const AvatarStage = lazy(() => import('./AvatarStage').then((module) => ({ default: module.AvatarStage })))

const languageOptions: { value: Language; label: string; short: string; caption: string; color: string }[] = [
  { value: 'golang', label: 'Golang', short: 'Go', caption: '并发 · Runtime · 工程化', color: '#6bd7e8' },
  { value: 'java', label: 'Java', short: 'J', caption: 'JVM · Spring · 分布式', color: '#ffb46e' },
  { value: 'python', label: 'Python', short: 'Py', caption: '语言 · 异步 · 服务端', color: '#e8d875' },
  { value: 'cpp', label: 'C++', short: 'C++', caption: '内存 · 并发 · 性能', color: '#ac9cff' },
  { value: 'agent-engineering', label: 'Agent 开发工程师', short: 'Agent', caption: 'Agent · RAG · 工具调用', color: '#59d4b5' },
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
  const [modelConfig, setModelConfig] = useState<ModelConfig>({ baseUrl: '', model: '', enabled: false, hasApiKey: false, speechAppId: '', speechResourceId: '', hasSpeechApiKey: false, ttsAppId: '', ttsResourceId: '', ttsSpeaker: '', ttsEnabled: false, hasTtsApiKey: false })
  const [catalog, setCatalog] = useState<SkillCatalog>()
  const [toast, setToast] = useState<Toast>()

  useEffect(() => {
    api.getModelConfig().then(setModelConfig).catch((error) => setToast({ type: 'error', message: error.message }))
    api.skills().then(setCatalog).catch((error) => setToast({ type: 'error', message: `Skill 加载失败：${error.message}` }))
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
        {screen === 'setup' && <Setup catalog={catalog} resume={resume} onResume={setResume} onStart={start} modelReady={modelConfig.enabled && modelConfig.hasApiKey} notify={setToast} />}
        {screen === 'interview' && session && (
          <InterviewScreen session={session} turns={turns} setTurns={setTurns} setSession={setSession} onFinish={finish} onBack={restart} notify={setToast} ttsEnabled={modelConfig.ttsEnabled && modelConfig.hasTtsApiKey && Boolean(modelConfig.ttsSpeaker)} />
        )}
        {screen === 'review' && report && <ReviewScreen report={report} onRestart={restart} />}
        {screen === 'learning' && <LearningPage catalog={catalog} notify={setToast} />}
        {screen === 'knowledge' && <KnowledgePage notify={setToast} onBack={restart} />}
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

function SkillSelect({ options, value, onChange, kind, placeholder }: { options: Skill[]; value: string; onChange: (value: string) => void; kind: 'domain' | 'interviewer'; placeholder: string }) {
  const [open, setOpen] = useState(false)
  const [query, setQuery] = useState('')
  const [activeIndex, setActiveIndex] = useState(-1)
  const rootRef = useRef<HTMLDivElement>(null)
  const searchRef = useRef<HTMLInputElement>(null)
  const selected = options.find((skill) => skill.id === value)
  const normalizedQuery = query.trim().toLocaleLowerCase()
  const filtered = options.filter((skill) => [skill.name, skill.shortLabel, skill.description, skill.id, ...(skill.topics ?? []), ...(skill.evaluationFocus ?? [])]
    .some((text) => text.toLocaleLowerCase().includes(normalizedQuery)))

  useEffect(() => {
    if (!open) return
    const close = (event: MouseEvent) => {
      if (!rootRef.current?.contains(event.target as Node)) setOpen(false)
    }
    document.addEventListener('mousedown', close)
    return () => document.removeEventListener('mousedown', close)
  }, [open])

  useEffect(() => {
    if (!open) return
    setQuery('')
    setActiveIndex(-1)
    window.requestAnimationFrame(() => searchRef.current?.focus())
  }, [open])

  const choose = (skill: Skill) => {
    onChange(skill.id)
    setOpen(false)
  }

  const handleSearchKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Escape') { setOpen(false); return }
    if (event.key === 'ArrowDown') { event.preventDefault(); setActiveIndex((index) => Math.min(index + 1, filtered.length - 1)); return }
    if (event.key === 'ArrowUp') { event.preventDefault(); setActiveIndex((index) => Math.max(index - 1, 0)); return }
    if (event.key === 'Enter' && filtered[activeIndex]) { event.preventDefault(); choose(filtered[activeIndex]) }
  }

  return (
    <div className={`skill-select ${open ? 'open' : ''}`} ref={rootRef} style={{ '--selected-skill-color': selected?.accent ?? '#6ee7c2' } as CSSProperties}>
      <button className="skill-select-trigger" type="button" onClick={() => setOpen((current) => !current)} aria-haspopup="listbox" aria-expanded={open}>
        {selected ? (
          <>
            <span className="skill-select-icon">{kind === 'domain' ? selected.shortLabel : <Bot size={18} />}</span>
            <span className="skill-select-copy"><strong>{selected.name}</strong><small>{kind === 'domain' ? selected.topics?.slice(0, 3).join(' · ') : selected.description}</small></span>
          </>
        ) : <span className="skill-select-placeholder">{placeholder}</span>}
        <ChevronDown className="skill-select-chevron" size={18} />
      </button>
      {open && (
        <div className="skill-select-menu">
          <label className="skill-search"><Search size={16} /><input ref={searchRef} value={query} onChange={(event) => { setQuery(event.target.value); setActiveIndex(0) }} onKeyDown={handleSearchKeyDown} placeholder="搜索名称、主题或关键字" aria-label="搜索 Skill" /></label>
          <div className="skill-options" role="listbox">
            {filtered.map((skill, index) => (
              <button key={skill.id} type="button" role="option" aria-selected={skill.id === value} className={`${skill.id === value ? 'selected' : ''} ${index === activeIndex ? 'active' : ''}`} onMouseEnter={() => setActiveIndex(index)} onClick={() => choose(skill)} style={{ '--skill-color': skill.accent } as CSSProperties}>
                <span className="skill-option-icon">{kind === 'domain' ? skill.shortLabel : <Bot size={17} />}</span>
                <span><strong>{skill.name}</strong><small>{kind === 'domain' ? skill.topics?.join(' · ') : skill.description}</small>{kind === 'interviewer' && <em>{skill.evaluationFocus?.join(' · ')}</em>}</span>
                {skill.id === value && <Check size={16} />}
              </button>
            ))}
            {!filtered.length && <div className="skill-empty"><Search size={20} /><strong>没有匹配的 Skill</strong><span>换一个关键字试试</span></div>}
          </div>
          <div className="skill-result-count">{normalizedQuery ? `找到 ${filtered.length} 个结果` : `共 ${options.length} 个 Skill`}</div>
        </div>
      )}
    </div>
  )
}

function Setup({ catalog, resume, onResume, onStart, modelReady, notify }: { catalog?: SkillCatalog; resume?: Resume; onResume: (value: Resume) => void; onStart: (value: Session) => void; modelReady: boolean; notify: (value: Toast) => void }) {
  const [candidateName, setCandidateName] = useState('')
  const [industry, setIndustry] = useState('computer')
  const [domainSkillId, setDomainSkillId] = useState('computer-golang')
  const [interviewerSkillId, setInterviewerSkillId] = useState('echo-coach')
  const [includeFoundation, setIncludeFoundation] = useState(true)
  const [videoEnabled, setVideoEnabled] = useState(false)
  const [speechLanguage, setSpeechLanguage] = useState('zh-CN')
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

  useEffect(() => {
    const available = (catalog?.domains ?? []).filter((skill) => skill.industry === industry && skill.language !== 'foundation')
    if (available.length && !available.some((skill) => skill.id === domainSkillId)) setDomainSkillId(available[0].id)
  }, [catalog, industry, domainSkillId])

  useEffect(() => {
    if (catalog?.interviewers.length && !catalog.interviewers.some((skill) => skill.id === interviewerSkillId)) setInterviewerSkillId(catalog.interviewers[0].id)
  }, [catalog, interviewerSkillId])

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
    try { onStart(await api.startInterview({ candidateName, resumeId: resume?.id, domainSkillId, interviewerSkillId, includeFoundation, videoEnabled, speechLanguage, difficulty, questionCount })) }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '面试创建失败' }) }
    finally { setStarting(false) }
  }

  const domainSkills = (catalog?.domains ?? []).filter((skill) => skill.industry === industry && skill.language !== 'foundation')

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
          <div className="section-heading"><span className="heading-icon violet"><Code2 size={19} /></span><div><h2>选择行业 / 领域 Skill</h2><p>每个 Skill 绑定独立知识库，后续可直接扩展新行业</p></div></div>
          <div className="industry-tabs">{(catalog?.industries ?? [{ id: 'computer', name: '计算机' }]).map((option) => <button key={option.id} className={industry === option.id ? 'selected' : ''} onClick={() => setIndustry(option.id)}>{option.name}</button>)}<span>更多行业 Skill 敬请期待</span></div>
          <SkillSelect options={domainSkills} value={domainSkillId} onChange={setDomainSkillId} kind="domain" placeholder="选择领域 Skill" />
        </section>

        <section className="panel interviewer-panel">
          <div className="section-heading"><span className="heading-icon mint"><Bot size={19} /></span><div><h2>选择面试官风格 Skill</h2><p>同一道题，不同面试官会用不同的评价重点和反馈语气</p></div></div>
          <SkillSelect options={catalog?.interviewers ?? []} value={interviewerSkillId} onChange={setInterviewerSkillId} kind="interviewer" placeholder="选择面试官 Skill" />
        </section>

        <section className="panel preferences-panel">
          <div className="section-heading"><span className="heading-icon amber"><Gauge size={19} /></span><div><h2>定一下节奏</h2><p>不用紧张，你可以在回答前慢慢组织思路</p></div></div>
          <label className="field-label">怎么称呼你</label>
          <div className="name-input"><UserRound size={17} /><input value={candidateName} maxLength={30} onChange={(event) => setCandidateName(event.target.value)} placeholder="候选人（可不填）" /></div>
          <label className="field-label">面试难度</label>
          <div className="difficulty-row">{difficultyOptions.map((option) => <button key={option.value} className={difficulty === option.value ? 'selected' : ''} onClick={() => setDifficulty(option.value)}><strong>{option.label}</strong><small>{option.caption}</small></button>)}</div>
          <label className="foundation-toggle"><input type="checkbox" checked={includeFoundation} onChange={(event) => setIncludeFoundation(event.target.checked)} /><span><LibraryBig size={16} /><strong>混入计算机基础公共库</strong><small>每场加入 1–2 道操作系统、网络、数据库或分布式基础题</small></span><i /></label>
          <label className="foundation-toggle"><input type="checkbox" checked={videoEnabled} onChange={(event) => setVideoEnabled(event.target.checked)} /><span>{videoEnabled ? <Video size={16} /> : <VideoOff size={16} />}<strong>开启视频面试</strong><small>进入面试后可选择摄像头；面试题、实时识别和朗读支持中文或英文</small></span><i /></label>
          <div className="count-row"><span><Mic size={16} />面试语言</span><div>{[{ value: 'zh-CN', label: '中文' }, { value: 'en-US', label: 'English' }].map((option) => <button key={option.value} className={speechLanguage === option.value ? 'selected' : ''} onClick={() => setSpeechLanguage(option.value)}>{option.label}</button>)}</div></div>
          {difficulty === 'mixed' && <div className="count-row"><span><MessageSquareText size={16} />题目数量</span><div>{[3, 5, 6].map((count) => <button key={count} className={questionCount === count ? 'selected' : ''} onClick={() => setQuestionCount(count)}>{count} 题</button>)}</div></div>}
          <button className="primary-action" onClick={start} disabled={starting || !catalog}>{starting ? <LoaderCircle className="spin" size={19} /> : <Sparkles size={18} />}开始模拟面试<ArrowRight size={18} /></button>
          <div className="start-hint"><span><Clock3 size={14} />约 {questionCount * 3}–{questionCount * 5} 分钟</span><span><Bot size={14} />{modelReady ? 'AI 深度点评' : '本地关键点评分'}</span></div>
        </section>

        <aside className="coach-card">
          <span className="coach-icon"><Lightbulb size={19} /></span><div><strong>一个小建议</strong><p>试着先讲结论，再讲原理，最后用项目经历收尾。清晰的结构往往比堆砌术语更有说服力。</p></div>
        </aside>
      </div>
    </div>
  )
}

function InterviewScreen({ session, turns, setTurns, setSession, onFinish, onBack, notify, ttsEnabled }: { session: Session; turns: Turn[]; setTurns: Dispatch<SetStateAction<Turn[]>>; setSession: (value: Session) => void; onFinish: (report: Report) => void; onBack: () => void; notify: (value: Toast) => void; ttsEnabled: boolean }) {
  const [answer, setAnswer] = useState('')
  const [submitting, setSubmitting] = useState(false)
  const [submitStage, setSubmitStage] = useState('')
  const [readyReport, setReadyReport] = useState<Report>()
  const [sideMessages, setSideMessages] = useState<IntentMessage[]>([])
  const [cameraStream, setCameraStream] = useState<MediaStream>()
  const [cameraError, setCameraError] = useState('')
  const [speechMode, setSpeechMode] = useState<SpeechMode>('idle')
  const [voiceMuted, setVoiceMuted] = useState(false)
  const [avatarSpeaking, setAvatarSpeaking] = useState(false)
  const [questionStartedAt, setQuestionStartedAt] = useState(Date.now())
  const [elapsed, setElapsed] = useState(0)
  const bottomRef = useRef<HTMLDivElement>(null)
  const videoRef = useRef<HTMLVideoElement>(null)
  const realtimeSpeechRef = useRef<RealtimeSpeechSession | null>(null)
  const speechConnectAbortRef = useRef<AbortController | null>(null)
  const speechModeRef = useRef<SpeechMode>('idle')
  const voicePlayerRef = useRef(new InterviewVoicePlayer({ onPlaybackChange: setAvatarSpeaking }))
  const lastSpokenRef = useRef('')
  const speechBaseAnswerRef = useRef('')
  const speechDraftRef = useRef('')
  useEffect(() => { const timer = window.setInterval(() => setElapsed(Math.floor((Date.now() - new Date(session.startedAt).getTime()) / 1000)), 1000); return () => clearInterval(timer) }, [session.startedAt])
  useEffect(() => { bottomRef.current?.scrollIntoView({ behavior: 'smooth' }) }, [turns, sideMessages, session.currentQuestion, readyReport])
  useEffect(() => {
    if (!session.videoEnabled) return
    let active = true
    let localStream: MediaStream | undefined
    navigator.mediaDevices?.getUserMedia({ video: true, audio: true }).then((stream) => {
      localStream = stream
      if (!active) { stream.getTracks().forEach((track) => track.stop()); return }
      setCameraStream(stream)
      if (videoRef.current) videoRef.current.srcObject = stream
    }).catch((error) => setCameraError(error instanceof Error ? error.message : '摄像头不可用'))
    return () => { active = false; localStream?.getTracks().forEach((track) => track.stop()) }
  }, [session.videoEnabled])
  useEffect(() => { if (videoRef.current && cameraStream) videoRef.current.srcObject = cameraStream }, [cameraStream])
  useEffect(() => { speechModeRef.current = speechMode }, [speechMode])
  useEffect(() => () => { speechConnectAbortRef.current?.abort(); realtimeSpeechRef.current?.stop() }, [])
  useEffect(() => () => voicePlayerRef.current.stop(), [])
  const currentAssistantMessage = sideMessages.filter((message) => message.role === 'assistant').at(-1)
  const currentVoice = currentAssistantMessage?.text || session.currentQuestion?.prompt
  const currentVoiceId = currentAssistantMessage?.id || session.currentQuestion?.id
  const avatarMode = deriveAvatarMode({ speechMode, speaking: avatarSpeaking, submitting })
  const speak = (text?: string, id?: string, force = false) => {
    if (speechModeRef.current !== 'idle' || !ttsEnabled || voiceMuted || !text || !id || (!force && lastSpokenRef.current === id)) return
    lastSpokenRef.current = id
    voicePlayerRef.current.speak(text).catch((error) => {
      if (error instanceof DOMException && error.name === 'AbortError') return
      notify({ type: 'error', message: error instanceof Error ? error.message : '面试官语音播放失败' })
    })
  }
  useEffect(() => {
    const opening = buildInterviewOpeningSpeech(session.speechLanguage, session.candidateName, session.interviewerOpening, session.currentQuestion?.prompt)
    const firstID = session.currentQuestion ? `opening:${session.currentQuestion.id}` : `opening:${session.id}`
    if (turns.length === 0 && sideMessages.length === 0) speak(opening, firstID)
    else speak(currentVoice, currentVoiceId)
  }, [currentVoiceId, currentVoice, session.currentQuestion, sideMessages.length, turns.length, ttsEnabled, voiceMuted, speechMode])

  const submit = async () => {
    const trimmed = answer.trim()
    if (!trimmed || submitting || speechMode !== 'idle' || !session.currentQuestion) return
    setSubmitting(true)
    setSubmitStage('理解你的意图')
    try {
      const question = session.currentQuestion
      window.setTimeout(() => setSubmitStage((stage) => stage ? '生成面试官反馈' : stage), 600)
      const result = await api.answer(session.id, trimmed, Math.floor((Date.now() - questionStartedAt) / 1000))
      setSubmitStage(result.accepted ? (result.requiresFollowUp ? '组织细化追问' : '更新题目进度') : '组织追问回复')
      setAnswer('')
      if (!result.accepted) {
        setSideMessages((previous) => [
          ...previous,
          { id: crypto.randomUUID(), role: 'candidate', text: trimmed, intent: result.intent || 'message' },
          ...(result.assistantReply ? [{ id: crypto.randomUUID(), role: 'assistant' as const, text: result.assistantReply, intent: result.intent || 'message' }] : []),
        ])
        return
      }
      const evaluation = result.evaluation
      if (!evaluation) throw new Error('本次回复缺少评价结果')
      setTurns((previous) => [...previous, { question, questionNumber: session.current + 1, followUp: result.requiresFollowUp, answer: trimmed, evaluation, sideMessages }])
      setSideMessages([])
      if (result.completed && result.report) {
        speak(buildCompletionSpeech(session.speechLanguage, evaluation.summary), `complete:${question.id}`)
        setReadyReport(result.report)
        setSession({ ...session, current: result.current, status: 'completed', currentQuestion: undefined })
      } else if (result.nextQuestion) {
        const transitionSpeech = result.requiresFollowUp
          ? buildFollowUpSpeech(session.speechLanguage, evaluation.summary, result.nextQuestion.prompt)
          : buildNextQuestionSpeech(session.speechLanguage, evaluation.summary, result.nextQuestion.prompt)
        speak(transitionSpeech, result.nextQuestion.id)
        setSession({ ...session, current: result.current, currentQuestion: result.nextQuestion })
        setQuestionStartedAt(Date.now())
      }
    } catch (error) {
      notify({ type: 'error', message: error instanceof Error ? error.message : '回答提交失败，请重试' })
    } finally { setSubmitting(false); setSubmitStage('') }
  }

  const toggleSpeech = async () => {
    if (speechMode === 'recording') {

      speechModeRef.current = 'finalizing'
      setSpeechMode('finalizing')
      realtimeSpeechRef.current?.stop()
      return
    }
    if (speechMode !== 'idle' || submitting) return
    voicePlayerRef.current.stop()
    speechBaseAnswerRef.current = answer.trimEnd()
    speechDraftRef.current = ''
    speechModeRef.current = 'connecting'
    setSpeechMode('connecting')
    const connectAbort = new AbortController()
    speechConnectAbortRef.current = connectAbort
    try {
      const speechSession = await startRealtimeSpeech(session.speechLanguage || 'zh-CN', (text, final) => {
        speechDraftRef.current = text.trim()
        setAnswer(mergeSpeechText(speechBaseAnswerRef.current, speechDraftRef.current))
        if (final) {
          realtimeSpeechRef.current = null
          speechModeRef.current = 'idle'
          setSpeechMode('idle')
        }
      }, (error) => {
        realtimeSpeechRef.current = null
        speechModeRef.current = 'idle'
        setSpeechMode('idle')
        notify({ type: 'error', message: error.message })
      }, connectAbort.signal)
      if (connectAbort.signal.aborted) {
        speechSession.stop()
        return
      }
      realtimeSpeechRef.current = speechSession
      speechConnectAbortRef.current = null
      speechModeRef.current = 'recording'
      setSpeechMode('recording')
    } catch (error) {
      speechConnectAbortRef.current = null
      speechModeRef.current = 'idle'
      setSpeechMode('idle')
      if (!(error instanceof DOMException && error.name === 'AbortError')) notify({ type: 'error', message: error instanceof Error ? error.message : '无法启动实时语音识别' })
    }
  }

  return (
    <div className="interview-page page-enter">
      <header className="interview-topbar">
        <button className="icon-button" onClick={onBack} title="退出本场面试"><ArrowLeft size={19} /></button>
        <div className="interview-title"><span className="live-dot" /><div><strong>{session.domainSkillName || languageLabel(session.language)}</strong><small>{session.candidateName} · {difficultyLabel(session.difficulty)} · {session.interviewerName}</small></div></div>
        <div className="progress-block"><div><span>进度</span><strong>{Math.min(session.current + 1, session.total)} / {session.total}</strong></div><div className="progress-track"><span style={{ width: `${Math.min(100, (session.current / session.total) * 100)}%` }} /></div></div>
        <div className="timer"><Clock3 size={16} />{formatTime(elapsed)}</div><button className={`voice-control ${voiceMuted || !ttsEnabled ? 'muted' : ''}`} type="button" title={!ttsEnabled ? '请先配置豆包 TTS' : voiceMuted ? '开启面试官语音' : '关闭面试官语音'} disabled={!ttsEnabled || speechMode !== 'idle'} onClick={() => { voicePlayerRef.current.stop(); setVoiceMuted((value) => !value) }}>{voiceMuted || !ttsEnabled ? <VolumeX size={17} /> : <Volume2 size={17} />}</button>
      </header>
      <div className="interview-layout">
        <section className="conversation">
          <div className="conversation-intro"><span><Bot size={21} /></span><div><strong>{session.interviewerName}</strong><p>{buildInterviewOpeningSpeech(session.speechLanguage, session.candidateName, session.interviewerOpening)}</p></div></div>
          {turns.map((turn) => <TurnCard key={turn.question.id} turn={turn} interviewerName={session.interviewerName} />)}
          {session.currentQuestion && (
            <div className="question-message">
              <div className="message-avatar"><Bot size={18} /></div>
              <div className="message-content"><div className="message-meta"><strong>{session.interviewerName}</strong><span>第 {session.current + 1} 题</span><button className="message-audio" type="button" title="重播题目" disabled={!ttsEnabled || voiceMuted || speechMode !== 'idle'} onClick={() => speak(session.currentQuestion?.prompt, `replay:${session.currentQuestion?.id}:${Date.now()}`, true)}><Volume2 size={14} /></button></div><p>{session.currentQuestion.prompt}</p><div className="question-tags">{session.currentQuestion.tags.map((tag) => <span key={tag}>{tag}</span>)}</div></div>
            </div>
          )}
          {sideMessages.map((message) => <IntentBubble key={message.id} message={message} interviewerName={session.interviewerName} />)}
          {readyReport && <div className="completion-card"><span><Sparkles size={26} /></span><h2>这场面试完成了</h2><p>你认真回答了 {readyReport.answered} 道题。标准答案、逐题评分和下一步练习建议都已经整理好。</p><button className="primary-action compact" onClick={() => onFinish(readyReport)}>查看我的复盘<ArrowRight size={18} /></button></div>}
          <div ref={bottomRef} />
        </section>
        <aside className="interview-aside">
          <Suspense fallback={<div className="avatar-stage"><div className="avatar-loading">正在唤醒面试官</div></div>}><AvatarStage mode={avatarMode} name={session.interviewerName} /></Suspense>
          {session.videoEnabled && <div className="video-card">{cameraStream ? <video ref={videoRef} autoPlay muted playsInline /> : <div><VideoOff size={21} /><p>{cameraError || '正在请求摄像头权限…'}</p></div>}<span>{isEnglishInterview(session.speechLanguage) ? 'English interview' : '中文面试'} · 火山实时语音识别</span></div>}
        </aside>
      </div>
      {session.currentQuestion && (
        <div className="composer-wrap"><div className="composer"><textarea autoFocus value={answer} maxLength={8000} onChange={(event) => setAnswer(event.target.value)} onKeyDown={(event) => { if ((event.metaKey || event.ctrlKey) && event.key === 'Enter') submit() }} placeholder="和面试官说说你的思路…" /><div className="composer-flow">{submitting && ['理解你的意图', '生成面试官反馈', '组织细化追问', '更新题目进度'].map((stage) => <span key={stage} className={submitStage === stage ? 'active' : ''}>{stage}</span>)}</div><div className="composer-footer"><span>{answer.length > 0 ? `${answer.length} 字` : '⌘ / Ctrl + Enter 发送'}</span><div><button className={`speech-button ${speechMode !== 'idle' ? 'active' : ''}`} onClick={toggleSpeech} type="button" disabled={submitting || speechMode === 'connecting' || speechMode === 'finalizing'}><Mic size={16} />{speechButtonText(speechMode)}</button><button onClick={submit} disabled={!answer.trim() || submitting || speechMode !== 'idle'}>{submitting ? <LoaderCircle className="spin" size={18} /> : <Send size={17} />}{submitting ? '处理中' : '发送'}</button></div></div></div></div>
      )}
    </div>
  )
}

function TurnCard({ turn, interviewerName }: { turn: Turn; interviewerName: string }) {
  const [open, setOpen] = useState(false)
  return (
    <div className="turn-block">
      <div className="question-message completed-question"><div className="message-avatar"><Bot size={18} /></div><div className="message-content"><div className="message-meta"><strong>{interviewerName}</strong><span>第 {turn.questionNumber} 题{turn.followUp ? '追问' : ''}</span></div><p>{turn.question.prompt}</p></div></div>
      {turn.sideMessages?.map((message) => <IntentBubble key={message.id} message={message} interviewerName={interviewerName} archived />)}
      <div className="answer-message"><div className="answer-bubble"><p>{turn.answer}</p></div><div className="user-avatar"><UserRound size={17} /></div></div>
      <button className={`inline-evaluation ${open ? 'open' : ''}`} onClick={() => setOpen(!open)}><span className={`score-dot score-${scoreBand(turn.evaluation.score)}`}>{turn.evaluation.score}</span><span><strong>{turn.evaluation.summary}</strong><small>{turn.evaluation.source === 'llm' ? 'Eino AI 点评' : '本地关键点评分'} · 点击{open ? '收起' : '展开'}</small></span><ChevronDown size={18} />
        {open && <div className="evaluation-details"><div><Check size={15} /><p>{turn.evaluation.strengths.join('；')}</p></div><div><Zap size={15} /><p>{turn.evaluation.improvements.join('；')}</p></div></div>}
      </button>
    </div>
  )
}

function IntentBubble({ message, interviewerName, archived = false }: { message: IntentMessage; interviewerName: string; archived?: boolean }) {
  if (message.role === 'candidate') {
    return <div className={`answer-message intent-message ${archived ? 'archived' : ''}`}><div className="answer-bubble"><p>{message.text}</p></div><div className="user-avatar"><UserRound size={17} /></div></div>
  }
  return <div className={`question-message intent-reply ${archived ? 'archived' : ''}`}><div className="message-avatar"><MessageSquareText size={17} /></div><div className="message-content"><div className="message-meta"><strong>{interviewerName}</strong><span>{intentLabel(message.intent)}</span></div><p>{message.text}</p></div></div>
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

export function KnowledgePage({ notify, onBack }: { notify: (value: Toast) => void; onBack: () => void }) {
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
    <header className="library-hero"><div><span className="eyebrow"><LibraryBig size={14} /> Knowledge Base</span><h1>每次出题，都让知识库<em>更好检索。</em></h1><p>语言库彼此独立，计算机基础作为公共库。每次面试实际出过的题都会累计出题次数并沉淀为 QA。</p></div><div className="library-hero-actions"><KnowledgeBackButton onBack={onBack} /><button className="secondary-action" onClick={() => setAdding(!adding)}>{adding ? <X size={17} /> : <Sparkles size={17} />}{adding ? '取消' : '新增 QA'}</button></div></header>
    <div className="knowledge-layout"><aside className="base-list"><span className="aside-label">知识库</span>{bases.map((base) => <button key={base.id} className={baseId === base.id ? 'selected' : ''} onClick={() => { setBaseId(base.id); setQuery('') }} style={{ '--base-color': base.accent } as CSSProperties}><span>{base.language === 'foundation' ? 'CS' : languageOptions.find((option) => option.value === base.language)?.short ?? base.language}</span><div><strong>{base.name}</strong><small>{base.itemCount} 个 QA · 已出题 {base.issuedCount} 次</small></div></button>)}</aside><section className="knowledge-main">
      {currentBase && <div className="base-summary"><span className="heading-icon violet"><LibraryBig size={19} /></span><div><h2>{currentBase.name}</h2><p>{currentBase.description}</p><div>{currentBase.topics.map((topic) => <span key={topic}>{topic}</span>)}</div></div></div>}
      {adding && <div className="qa-form"><label>问题<input value={question} onChange={(event) => setQuestion(event.target.value)} placeholder="输入希望沉淀的问题" /></label><label>标准答案<textarea value={answer} onChange={(event) => setAnswer(event.target.value)} placeholder="输入可复盘、可检索的标准答案" /></label><button className="primary-action compact" disabled={!question.trim() || !answer.trim()} onClick={addQA}><Check size={16} />写入知识库</button></div>}
      <div className="knowledge-search"><div><SearchIcon /><input value={query} onChange={(event) => setQuery(event.target.value)} onKeyDown={(event) => { if (event.key === 'Enter') search() }} placeholder="检索题目、答案或标签…" /></div><button onClick={() => search()}>检索</button></div>
      {loading ? <div className="empty-state compact"><LoaderCircle className="spin" size={23} /></div> : <div className="qa-list">{items.map((item) => <article key={item.id} className={openItem === item.id ? 'open' : ''}><button onClick={() => setOpenItem(openItem === item.id ? undefined : item.id)}><span className="qa-source">{item.source === 'manual' ? '手动' : item.source === 'interview' ? '面试沉淀' : '内置'}</span><div><h3>{item.question}</h3><small>{item.tags.join(' · ')}{item.issuedCount > 0 && ` · 已出题 ${item.issuedCount} 次`}</small></div><ChevronDown size={18} /></button>{openItem === item.id && <div className="qa-answer"><span>标准答案</span><p>{item.answer}</p><div>{item.keyPoints.map((point) => <em key={point}>{point.split('/')[0]}</em>)}</div></div>}</article>)}</div>}
    </section></div>
  </div>
}

export function KnowledgeBackButton({ onBack }: { onBack: () => void }) {
  return <button className="secondary-action" onClick={onBack}><ArrowLeft size={17} />返回面试准备</button>
}

function SearchIcon() { return <Target size={16} /> }
function resourceKindLabel(kind: string) { return ({ all: '全部类型', article: '文章', video: '视频', course: '课程', docs: '文档' } as Record<string, string>)[kind] ?? kind }
function resourceKindIcon(kind: string) { if (kind === 'video') return <PlayCircle size={13} />; if (kind === 'article') return <Newspaper size={13} />; return <BookOpen size={13} /> }
function resourceDate(resource: LearningResource) { const date = resource.publishedAt ? new Date(resource.publishedAt) : undefined; return date && date.getFullYear() > 1900 ? date.toLocaleDateString('zh-CN') : resource.source }
function intentLabel(intent: string) { return ({ hint: '提示', clarify: '换个说法', repeat: '重复题目', skip: '跳过', off_topic: '拉回题目', smalltalk: '闲聊回应', message: '继续对话' } as Record<string, string>)[intent] ?? '继续对话' }
function speechButtonText(mode: SpeechMode) { return ({ idle: '语音', connecting: '连接中', recording: '停止录音', finalizing: '整理文字' } as Record<SpeechMode, string>)[mode] }
function mergeSpeechText(base: string, transcript: string) { return [base.trimEnd(), transcript.trim()].filter(Boolean).join(base.trim() ? '\n' : '') }

function SettingsDrawer({ open, value, onClose, onSaved, notify }: { open: boolean; value: ModelConfig; onClose: () => void; onSaved: (value: ModelConfig) => void; notify: (value: Toast) => void }) {
  const [apiKey, setApiKey] = useState('')
  const [baseUrl, setBaseUrl] = useState(value.baseUrl)
  const [model, setModel] = useState(value.model)
  const [speechApiKey, setSpeechApiKey] = useState('')
  const [speechAppId, setSpeechAppId] = useState(value.speechAppId)
  const [speechResourceId, setSpeechResourceId] = useState(value.speechResourceId || 'volc.bigasr.sauc.duration')
  const [ttsApiKey, setTtsApiKey] = useState('')
  const [ttsAppId, setTtsAppId] = useState(value.ttsAppId)
  const [ttsResourceId, setTtsResourceId] = useState(value.ttsResourceId || 'seed-tts-2.0')
  const [ttsSpeaker, setTtsSpeaker] = useState(value.ttsSpeaker)
  const [ttsEnabled, setTtsEnabled] = useState(value.ttsEnabled)
  const [enabled, setEnabled] = useState(value.enabled)
  const [saving, setSaving] = useState(false)
  const [testing, setTesting] = useState(false)
  useEffect(() => { if (open) { setApiKey(''); setBaseUrl(value.baseUrl); setModel(value.model); setSpeechApiKey(''); setSpeechAppId(value.speechAppId); setSpeechResourceId(value.speechResourceId || 'volc.bigasr.sauc.duration'); setTtsApiKey(''); setTtsAppId(value.ttsAppId); setTtsResourceId(value.ttsResourceId || 'seed-tts-2.0'); setTtsSpeaker(value.ttsSpeaker); setTtsEnabled(value.ttsEnabled); setEnabled(value.enabled) } }, [open, value])
  if (!open) return null
  const save = async () => {
    setSaving(true)
    try { const next = await api.saveModelConfig({ apiKey, baseUrl, model, enabled, speechApiKey, speechAppId, speechResourceId, ttsApiKey, ttsAppId, ttsResourceId, ttsSpeaker, ttsEnabled }); onSaved(next); notify({ type: 'success', message: '模型与双向语音配置已保存' }); onClose() }
    catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '保存失败' }) }
    finally { setSaving(false) }
  }
  const test = async () => {
    setTesting(true)
    try {
      const next = await api.saveModelConfig({ apiKey, baseUrl, model, enabled: true, speechApiKey, speechAppId, speechResourceId, ttsApiKey, ttsAppId, ttsResourceId, ttsSpeaker, ttsEnabled }); onSaved(next)
      const result = await api.testModelConfig(); setEnabled(true); notify({ type: 'success', message: result.message })
    } catch (error) { notify({ type: 'error', message: error instanceof Error ? error.message : '连接失败' }) }
    finally { setTesting(false) }
  }
  return <div className="drawer-backdrop" onMouseDown={(event) => { if (event.currentTarget === event.target) onClose() }}><aside className="settings-drawer"><header><div><span className="heading-icon violet"><Settings2 size={19} /></span><section><h2>模型与双向语音</h2><p>对话模型 + 豆包 ASR/TTS</p></section></div><button className="icon-button" onClick={onClose}><X size={19} /></button></header><div className="drawer-content"><div className="settings-callout"><ShieldCheck size={19} /><p><strong>凭据保存在本机</strong><br />所有 Key 均以 0600 权限写入应用数据目录，不会进入项目代码或日志。</p></div><label className="setting-field"><span>对话 API Key {value.hasApiKey && <em>已保存</em>}</span><input type="password" autoComplete="off" value={apiKey} onChange={(event) => setApiKey(event.target.value)} placeholder={value.hasApiKey ? '留空以继续使用已保存的 Key' : 'sk-...'} /></label><label className="setting-field"><span>Base URL <small>可留空使用默认地址</small></span><input value={baseUrl} onChange={(event) => setBaseUrl(event.target.value)} placeholder="https://api.example.com/v1" /></label><label className="setting-field"><span>Model</span><input value={model} onChange={(event) => setModel(event.target.value)} placeholder="对话 / 出题 / 点评模型 ID" /></label><div className="settings-note"><Mic size={16} /><p><strong>豆包实时语音识别（ASR）</strong><br />旧版控制台填写 App ID + Access Token；新版控制台只填写 API Key，App ID 留空。</p></div><label className="setting-field"><span>ASR Access Token / API Key {value.hasSpeechApiKey && <em>已保存</em>}</span><input type="password" autoComplete="off" value={speechApiKey} onChange={(event) => setSpeechApiKey(event.target.value)} placeholder={value.hasSpeechApiKey ? '留空以继续使用已保存的凭据' : '旧版 Access Token 或新版 API Key'} /></label><label className="setting-field"><span>ASR App ID <small>旧版控制台必填</small></span><input value={speechAppId} onChange={(event) => setSpeechAppId(event.target.value)} placeholder="旧版火山语音应用 App ID" /></label><label className="setting-field"><span>ASR Resource ID</span><input value={speechResourceId} onChange={(event) => setSpeechResourceId(event.target.value)} placeholder="volc.bigasr.sauc.duration" /></label><div className="settings-note"><Volume2 size={16} /><p><strong>豆包语音合成（TTS）</strong><br />TTS 需要独立开通资源并选择控制台音色，不能复用 ASR Resource ID。</p></div><label className="setting-field"><span>TTS Access Token / API Key {value.hasTtsApiKey && <em>已保存</em>}</span><input type="password" autoComplete="off" value={ttsApiKey} onChange={(event) => setTtsApiKey(event.target.value)} placeholder={value.hasTtsApiKey ? '留空以继续使用已保存的凭据' : 'TTS Access Token 或 API Key'} /></label><label className="setting-field"><span>TTS App ID <small>旧版控制台填写</small></span><input value={ttsAppId} onChange={(event) => setTtsAppId(event.target.value)} placeholder="旧版 TTS App ID" /></label><label className="setting-field"><span>TTS Resource ID</span><input value={ttsResourceId} onChange={(event) => setTtsResourceId(event.target.value)} placeholder="seed-tts-2.0" /></label><label className="setting-field"><span>音色 ID</span><input value={ttsSpeaker} onChange={(event) => setTtsSpeaker(event.target.value)} placeholder="从豆包语音控制台音色列表复制" /></label><label className="toggle-row"><span><strong>自动朗读面试官内容</strong><small>朗读开场、题目、追问和反馈</small></span><input type="checkbox" checked={ttsEnabled} onChange={(event) => setTtsEnabled(event.target.checked)} /><i /></label><label className="toggle-row"><span><strong>启用 AI 深度点评</strong><small>不可用时会自动降级为本地评分</small></span><input type="checkbox" checked={enabled} onChange={(event) => setEnabled(event.target.checked)} /><i /></label></div><footer><button className="secondary-action" onClick={test} disabled={testing || (!apiKey && !value.hasApiKey) || !model}>{testing ? <LoaderCircle className="spin" size={16} /> : <Zap size={16} />}测试对话模型</button><button className="primary-action compact" onClick={save} disabled={saving}>{saving ? <LoaderCircle className="spin" size={16} /> : <Check size={16} />}保存设置</button></footer></aside></div>
}

function languageLabel(language: Language) { return languageOptions.find((option) => option.value === language)?.label ?? language }
function difficultyLabel(value: Difficulty) { return difficultyOptions.find((option) => option.value === value)?.label ?? value }
function formatTime(seconds: number) { const minutes = Math.floor(seconds / 60); return `${String(minutes).padStart(2, '0')}:${String(seconds % 60).padStart(2, '0')}` }
function scoreBand(score: number) { return score >= 80 ? 'great' : score >= 60 ? 'good' : 'focus' }
function performanceText(score: number) { return score >= 80 ? '表现出色' : score >= 60 ? '稳步进阶' : '继续加油' }

export default App
