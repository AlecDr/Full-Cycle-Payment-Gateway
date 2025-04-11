import type React from "react"
import { Header } from "@/components/header"

export default function AuthenticatedLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <>
      <Header />
      <main className="container mx-auto px-4 py-8">{children}</main>
      <footer className="container mx-auto px-4 py-4 text-center text-slate-500 text-sm">
        © 2025 Full Cycle Gateway. Todos os direitos reservados.
      </footer>
    </>
  )
}
