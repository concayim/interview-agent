import { useEffect, useRef, useState } from 'react'
import * as THREE from 'three'
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js'
import { VRM, VRMHumanBoneName, VRMLoaderPlugin, VRMUtils } from '@pixiv/three-vrm'
import type { AvatarMode } from './avatar-state'

const modeLabels: Record<AvatarMode, string> = {
  idle: '等待回答',
  listening: '正在聆听',
  thinking: '正在思考',
  speaking: '正在说话',
}

export function AvatarStage({ mode, name }: { mode: AvatarMode; name: string }) {
  const hostRef = useRef<HTMLDivElement>(null)
  const modeRef = useRef(mode)
  const [loaded, setLoaded] = useState(false)
  const [failed, setFailed] = useState(false)

  useEffect(() => { modeRef.current = mode }, [mode])
  useEffect(() => {
    const host = hostRef.current
    if (!host) return
    let disposed = false
    let frame = 0
    let vrm: VRM | undefined
    const scene = new THREE.Scene()
    const camera = new THREE.PerspectiveCamera(24, 1, 0.1, 20)
    camera.position.set(0, 1.35, 3.35)
    camera.lookAt(0, 1.25, 0)
    const renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true, powerPreference: 'high-performance' })
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))
    renderer.outputColorSpace = THREE.SRGBColorSpace
    renderer.setClearColor(0x000000, 0)
    host.appendChild(renderer.domElement)
    scene.add(new THREE.HemisphereLight(0xe9f5ff, 0x19223a, 2.4))
    const keyLight = new THREE.DirectionalLight(0xffffff, 2.8)
    keyLight.position.set(1.6, 2.5, 2.4)
    scene.add(keyLight)
    const rimLight = new THREE.DirectionalLight(0x78e2cb, 1.8)
    rimLight.position.set(-2, 1.4, -1)
    scene.add(rimLight)

    const resize = () => {
      const { width, height } = host.getBoundingClientRect()
      if (!width || !height) return
      renderer.setSize(width, height, false)
      camera.aspect = width / height
      camera.updateProjectionMatrix()
    }
    const observer = new ResizeObserver(resize)
    observer.observe(host)
    resize()

    const loader = new GLTFLoader()
    loader.register((parser) => new VRMLoaderPlugin(parser))
    loader.load('./models/interviewer.vrm', (gltf) => {
      if (disposed) return
      vrm = gltf.userData.vrm as VRM
      VRMUtils.removeUnnecessaryVertices(gltf.scene)
      VRMUtils.combineSkeletons(gltf.scene)
      vrm.scene.rotation.y = Math.PI
      vrm.scene.position.y = 0.02
      const leftUpperArm = vrm.humanoid?.getNormalizedBoneNode(VRMHumanBoneName.LeftUpperArm)
      const rightUpperArm = vrm.humanoid?.getNormalizedBoneNode(VRMHumanBoneName.RightUpperArm)
      const leftLowerArm = vrm.humanoid?.getNormalizedBoneNode(VRMHumanBoneName.LeftLowerArm)
      const rightLowerArm = vrm.humanoid?.getNormalizedBoneNode(VRMHumanBoneName.RightLowerArm)
      if (leftUpperArm) leftUpperArm.rotation.z = -1.08
      if (rightUpperArm) rightUpperArm.rotation.z = 1.08
      if (leftLowerArm) leftLowerArm.rotation.y = -0.12
      if (rightLowerArm) rightLowerArm.rotation.y = 0.12
      scene.add(vrm.scene)
      setLoaded(true)
    }, undefined, () => { if (!disposed) setFailed(true) })

    const clock = new THREE.Clock()
    const animate = () => {
      frame = requestAnimationFrame(animate)
      const delta = Math.min(clock.getDelta(), 0.05)
      const t = clock.elapsedTime
      if (vrm) {
        const avatarMode = modeRef.current
        const head = vrm.humanoid?.getNormalizedBoneNode(VRMHumanBoneName.Head)
        const chest = vrm.humanoid?.getNormalizedBoneNode(VRMHumanBoneName.Chest)
        if (head) {
          head.rotation.y = Math.sin(t * 0.55) * 0.035
          head.rotation.x = avatarMode === 'thinking' ? -0.055 + Math.sin(t) * 0.012 : Math.sin(t * 0.7) * 0.012
        }
        if (chest) chest.rotation.z = Math.sin(t * 0.72) * 0.009
        const blink = Math.max(0, (Math.sin(t * 1.15) - 0.985) * 65)
        vrm.expressionManager?.setValue('blink', blink)
        vrm.expressionManager?.setValue('happy', avatarMode === 'speaking' ? 0.18 : avatarMode === 'listening' ? 0.08 : 0.03)
        vrm.expressionManager?.setValue('aa', avatarMode === 'speaking' ? 0.18 + Math.abs(Math.sin(t * 8.5)) * 0.52 : 0)
        vrm.expressionManager?.setValue('oh', avatarMode === 'speaking' ? Math.abs(Math.cos(t * 6.4)) * 0.16 : 0)
        vrm.update(delta)
      }
      renderer.render(scene, camera)
    }
    animate()

    return () => {
      disposed = true
      cancelAnimationFrame(frame)
      observer.disconnect()
      vrm && VRMUtils.deepDispose(vrm.scene)
      renderer.dispose()
      renderer.domElement.remove()
    }
  }, [])

  return (
    <section className={`avatar-stage mode-${mode}`} aria-label={`${name} AI 形象`}>
      <div className="avatar-canvas" ref={hostRef} />
      {!loaded && !failed && <div className="avatar-loading">正在唤醒面试官</div>}
      {failed && <div className="avatar-loading">形象加载失败，文字问答仍可继续</div>}
      <div className="avatar-status"><span /><div><strong>{name}</strong><small>{modeLabels[mode]}</small></div></div>
      <a href="https://github.com/vrm-c/vrm-specification/tree/master/samples/Seed-san" target="_blank" rel="noreferrer">Seed-san by VirtualCast, Inc.</a>
    </section>
  )
}
