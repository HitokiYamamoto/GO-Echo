import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    host: true, // 0.0.0.0 にバインドする
    port: 5173, // 必要に応じてポート固定
    strictPort: true, // ポート競合時にエラーを出す（オプション）
  },
});
