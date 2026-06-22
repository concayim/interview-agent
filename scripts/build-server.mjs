import { mkdir } from 'node:fs/promises'
import { spawnSync } from 'node:child_process'
import { join } from 'node:path'

const outputDir = join(process.cwd(), 'resources', 'bin')
await mkdir(outputDir, { recursive: true })
const binary = process.platform === 'win32' ? 'interview-agent-server.exe' : 'interview-agent-server'
const result = spawnSync('go', ['build', '-trimpath', '-o', join(outputDir, binary), './cmd/server'], {
  stdio: 'inherit',
  env: process.env,
})
if (result.status !== 0) process.exit(result.status ?? 1)
