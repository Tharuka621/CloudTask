import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1', // explicitly use IPv4 for Docker compatibility
        changeOrigin: true,
      },
      '/ws': {
        target: 'ws://127.0.0.1',
        ws: true,
      }
    }
  }
})
