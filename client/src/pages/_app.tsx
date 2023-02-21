import "./globals.css"

import { Poppins } from "@next/font/google"
import type { AppProps } from "next/app"

const poppins = Poppins({ subsets: ["latin"], weight: ["200", "400", "700"] })

export default function App({ Component, pageProps }: AppProps) {
  return (
    <main className={poppins.className}>
      <Component {...pageProps} />
    </main>
  )
}
