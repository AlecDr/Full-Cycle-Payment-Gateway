"use client"

import { useRouter } from "next/navigation"
import { ArrowLeft, Download, CheckCircle } from "lucide-react"
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"

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

interface InvoiceDetailsPageProps {
  params: {
    id: string
  }
}

export default function InvoiceDetailsPage({ params }: InvoiceDetailsPageProps) {
  const router = useRouter()
  const invoiceId = `#${params.id}`

  // Mock data for the invoice details
  const invoice = {
    id: "#INV-001",
    status: "approved" as const,
    createdAt: "30/03/2025 às 14:30",
    value: "R$ 1.500,00",
    creationDate: "30/03/2025 14:30",
    lastUpdate: "30/03/2025 14:35",
    description: "Compra Online #123",
    paymentMethod: "Cartão de Crédito",
    lastDigits: "1234",
    cardHolder: "João da Silva",
    accountId: "ACC-12345",
    clientIp: "192.168.1.1",
    device: "Desktop - Chrome",
    transactionSteps: [
      {
        title: "Fatura Criada",
        date: "30/03/2025 14:30",
      },
      {
        title: "Pagamento Processado",
        date: "30/03/2025 14:32",
      },
      {
        title: "Transação Aprovada",
        date: "30/03/2025 14:35",
      },
    ],
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Button variant="outline" size="icon" onClick={() => router.push("/invoices")}>
            <ArrowLeft size={18} />
          </Button>
          <div>
            <h1 className="text-2xl font-bold flex items-center gap-3 text-slate-800">
              Fatura {invoiceId}
              <StatusBadge status="approved" />
            </h1>
            <p className="text-slate-500 text-sm">Criada em {invoice.createdAt}</p>
          </div>
        </div>
        <Button className="flex items-center gap-2">
          <Download size={16} />
          Download PDF
        </Button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <h2 className="text-xl font-semibold text-slate-800">Informações da Fatura</h2>
          </CardHeader>
          <CardContent>
            <dl className="space-y-4">
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">ID da Fatura</dt>
                <dd className="text-slate-800">{invoice.id}</dd>
              </div>
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">Valor</dt>
                <dd className="text-slate-800">{invoice.value}</dd>
              </div>
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">Data de Criação</dt>
                <dd className="text-slate-800">{invoice.creationDate}</dd>
              </div>
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">Última Atualização</dt>
                <dd className="text-slate-800">{invoice.lastUpdate}</dd>
              </div>
              <div className="flex justify-between py-2">
                <dt className="text-slate-500">Descrição</dt>
                <dd className="text-slate-800">{invoice.description}</dd>
              </div>
            </dl>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <h2 className="text-xl font-semibold text-slate-800">Status da Transação</h2>
          </CardHeader>
          <CardContent>
            <div className="space-y-6">
              {invoice.transactionSteps.map((step, index) => (
                <div key={index} className="flex gap-4">
                  <div className="flex-shrink-0 mt-1">
                    <CheckCircle className="h-6 w-6 text-green-500" />
                  </div>
                  <div className="space-y-1">
                    <p className="font-medium text-slate-800">{step.title}</p>
                    <p className="text-sm text-slate-500">{step.date}</p>
                  </div>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <h2 className="text-xl font-semibold text-slate-800">Método de Pagamento</h2>
          </CardHeader>
          <CardContent>
            <dl className="space-y-4">
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">Tipo</dt>
                <dd className="text-slate-800">{invoice.paymentMethod}</dd>
              </div>
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">Últimos Dígitos</dt>
                <dd className="text-slate-800">**** **** **** {invoice.lastDigits}</dd>
              </div>
              <div className="flex justify-between py-2">
                <dt className="text-slate-500">Titular</dt>
                <dd className="text-slate-800">{invoice.cardHolder}</dd>
              </div>
            </dl>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <h2 className="text-xl font-semibold text-slate-800">Dados Adicionais</h2>
          </CardHeader>
          <CardContent>
            <dl className="space-y-4">
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">ID da Conta</dt>
                <dd className="text-slate-800">{invoice.accountId}</dd>
              </div>
              <div className="flex justify-between py-2 border-b border-gray-200">
                <dt className="text-slate-500">IP do Cliente</dt>
                <dd className="text-slate-800">{invoice.clientIp}</dd>
              </div>
              <div className="flex justify-between py-2">
                <dt className="text-slate-500">Dispositivo</dt>
                <dd className="text-slate-800">{invoice.device}</dd>
              </div>
            </dl>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
