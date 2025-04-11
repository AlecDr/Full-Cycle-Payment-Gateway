"use client"

import { useState } from "react"
import Link from "next/link"
import { Eye, Download, Plus, ChevronLeft, ChevronRight } from "lucide-react"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { cn } from "@/lib/utils"
import { formatCurrency } from "@/lib/utils"

// Mock data for invoices
const invoices = [
  {
    id: "#INV-001",
    date: "30/03/2025",
    description: "Compra Online #123",
    value: 1500.0,
    status: "approved" as const,
  },
  {
    id: "#INV-002",
    date: "29/03/2025",
    description: "Serviço Premium",
    value: 15000.0,
    status: "pending" as const,
  },
  {
    id: "#INV-003",
    date: "28/03/2025",
    description: "Assinatura Mensal",
    value: 99.9,
    status: "rejected" as const,
  },
]

// Inline status badge component
function StatusBadge({ status }: { status: "approved" | "pending" | "rejected" }) {
  return (
    <span
      className={cn("px-3 py-1 rounded-full text-xs font-medium", {
        "bg-green-100 text-green-800": status === "approved",
        "bg-yellow-100 text-yellow-800": status === "pending",
        "bg-red-100 text-red-800": status === "rejected",
      })}
    >
      {status === "approved" && "Aprovado"}
      {status === "pending" && "Pendente"}
      {status === "rejected" && "Rejeitado"}
    </span>
  )
}

export default function InvoicesPage() {
  const [searchTerm, setSearchTerm] = useState("")
  const [currentPage, setCurrentPage] = useState(1)
  const totalResults = 50
  const resultsPerPage = 10
  const totalPages = Math.ceil(totalResults / resultsPerPage)

  return (
    <Card>
      <CardHeader>
        <div className="flex justify-between items-start">
          <div>
            <CardTitle>Faturas</CardTitle>
            <CardDescription>Gerencie suas faturas e acompanhe os pagamentos</CardDescription>
          </div>
          <Link href="/invoices/new">
            <Button className="flex items-center gap-2">
              <Plus size={16} />
              Nova Fatura
            </Button>
          </Link>
        </div>
      </CardHeader>
      <CardContent>
        <div className="bg-gray-50 border border-gray-200 rounded-lg p-4 mb-6">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            <div>
              <label htmlFor="status" className="block text-sm font-medium mb-1 text-slate-700">
                Status
              </label>
              <select id="status" className="w-full h-10 rounded-md border border-gray-300 bg-white px-3 py-2 text-sm">
                <option value="all">Todos</option>
                <option value="approved">Aprovado</option>
                <option value="pending">Pendente</option>
                <option value="rejected">Rejeitado</option>
              </select>
            </div>
            <div>
              <label htmlFor="startDate" className="block text-sm font-medium mb-1 text-slate-700">
                Data Inicial
              </label>
              <Input id="startDate" type="text" placeholder="dd/mm/aaaa" />
            </div>
            <div>
              <label htmlFor="endDate" className="block text-sm font-medium mb-1 text-slate-700">
                Data Final
              </label>
              <Input id="endDate" type="text" placeholder="dd/mm/aaaa" />
            </div>
            <div>
              <label htmlFor="search" className="block text-sm font-medium mb-1 text-slate-700">
                Buscar
              </label>
              <Input
                id="search"
                type="text"
                placeholder="ID ou descrição"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
              />
            </div>
          </div>
        </div>

        <div className="overflow-x-auto">
          <table className="w-full">
            <thead>
              <tr className="border-b border-gray-200">
                <th className="text-left py-3 px-4 font-medium text-sm text-slate-700">ID</th>
                <th className="text-left py-3 px-4 font-medium text-sm text-slate-700">DATA</th>
                <th className="text-left py-3 px-4 font-medium text-sm text-slate-700">DESCRIÇÃO</th>
                <th className="text-left py-3 px-4 font-medium text-sm text-slate-700">VALOR</th>
                <th className="text-left py-3 px-4 font-medium text-sm text-slate-700">STATUS</th>
                <th className="text-left py-3 px-4 font-medium text-sm text-slate-700">AÇÕES</th>
              </tr>
            </thead>
            <tbody>
              {invoices.map((invoice) => (
                <tr key={invoice.id} className="border-b border-gray-200">
                  <td className="py-4 px-4 text-slate-800">{invoice.id}</td>
                  <td className="py-4 px-4 text-slate-800">{invoice.date}</td>
                  <td className="py-4 px-4 text-slate-800">{invoice.description}</td>
                  <td className="py-4 px-4 text-slate-800">{formatCurrency(invoice.value)}</td>
                  <td className="py-4 px-4">
                    <StatusBadge status={invoice.status} />
                  </td>
                  <td className="py-4 px-4">
                    <div className="flex gap-2">
                      <Link href={`/invoices/${invoice.id.replace("#", "")}`}>
                        <Button variant="ghost" size="icon">
                          <Eye size={18} />
                        </Button>
                      </Link>
                      <Button variant="ghost" size="icon">
                        <Download size={18} />
                      </Button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        <div className="flex items-center justify-between mt-6">
          <div className="text-sm text-slate-500">Mostrando 1 - 3 de 50 resultados</div>
          <div className="flex items-center gap-1">
            <Button
              variant="outline"
              size="icon"
              disabled={currentPage === 1}
              onClick={() => setCurrentPage(currentPage - 1)}
            >
              <ChevronLeft size={16} />
            </Button>
            {[1, 2, 3].map((page) => (
              <Button
                key={page}
                variant={currentPage === page ? "default" : "outline"}
                size="sm"
                onClick={() => setCurrentPage(page)}
              >
                {page}
              </Button>
            ))}
            <Button
              variant="outline"
              size="icon"
              disabled={currentPage === totalPages}
              onClick={() => setCurrentPage(currentPage + 1)}
            >
              <ChevronRight size={16} />
            </Button>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
