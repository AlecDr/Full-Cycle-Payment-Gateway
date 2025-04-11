"use client"

import type React from "react"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { LockIcon } from "lucide-react"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { formatCurrency } from "@/lib/utils"

export default function NewInvoicePage() {
  const router = useRouter()
  const [amount, setAmount] = useState("")
  const [description, setDescription] = useState("")
  const [cardNumber, setCardNumber] = useState("")
  const [expiryDate, setExpiryDate] = useState("")
  const [cvv, setCvv] = useState("")
  const [cardName, setCardName] = useState("")

  const numericAmount = Number.parseFloat(amount.replace(/[^\d,]/g, "").replace(",", ".")) || 0
  const fee = numericAmount * 0.02
  const total = numericAmount + fee

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // In a real app, we would submit the form data to an API
    router.push("/invoices")
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Criar Nova Fatura</CardTitle>
        <CardDescription>Preencha os dados abaixo para processar um novo pagamento</CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="space-y-8">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="space-y-6">
              <div className="space-y-2">
                <label htmlFor="amount" className="block text-sm font-medium text-slate-700">
                  Valor
                </label>
                <Input
                  id="amount"
                  type="text"
                  placeholder="R$ 0,00"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                  required
                />
              </div>

              <div className="space-y-2">
                <label htmlFor="description" className="block text-sm font-medium text-slate-700">
                  Descrição
                </label>
                <Textarea
                  id="description"
                  placeholder="Descreva o motivo do pagamento"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  required
                />
              </div>
            </div>

            <div className="space-y-6 bg-gray-50 border border-gray-200 rounded-lg p-6">
              <h3 className="text-lg font-medium text-slate-800">Dados do Cartão</h3>

              <div className="space-y-2">
                <label htmlFor="cardNumber" className="block text-sm font-medium text-slate-700">
                  Número do Cartão
                </label>
                <div className="relative">
                  <Input
                    id="cardNumber"
                    type="text"
                    placeholder="0000 0000 0000 0000"
                    value={cardNumber}
                    onChange={(e) => setCardNumber(e.target.value)}
                    required
                  />
                  <div className="absolute right-3 top-1/2 transform -translate-y-1/2">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <rect x="2" y="5" width="20" height="14" rx="2" stroke="#9CA3AF" strokeWidth="2" />
                    </svg>
                  </div>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <label htmlFor="expiryDate" className="block text-sm font-medium text-slate-700">
                    Data de Expiração
                  </label>
                  <Input
                    id="expiryDate"
                    type="text"
                    placeholder="MM/AA"
                    value={expiryDate}
                    onChange={(e) => setExpiryDate(e.target.value)}
                    required
                  />
                </div>

                <div className="space-y-2">
                  <label htmlFor="cvv" className="block text-sm font-medium text-slate-700">
                    CVV
                  </label>
                  <Input
                    id="cvv"
                    type="text"
                    placeholder="123"
                    value={cvv}
                    onChange={(e) => setCvv(e.target.value)}
                    required
                  />
                </div>
              </div>

              <div className="space-y-2">
                <label htmlFor="cardName" className="block text-sm font-medium text-slate-700">
                  Nome no Cartão
                </label>
                <Input
                  id="cardName"
                  type="text"
                  placeholder="Como aparece no cartão"
                  value={cardName}
                  onChange={(e) => setCardName(e.target.value)}
                  required
                />
              </div>
            </div>
          </div>

          <div className="bg-gray-50 border border-gray-200 rounded-lg p-6">
            <div className="space-y-4">
              <div className="flex justify-between">
                <span className="text-slate-600">Subtotal</span>
                <span className="text-slate-800">{formatCurrency(numericAmount)}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-slate-600">Taxa de Processamento (2%)</span>
                <span className="text-slate-800">{formatCurrency(fee)}</span>
              </div>
              <div className="border-t border-gray-200 pt-4 flex justify-between">
                <span className="font-medium text-slate-800">Total</span>
                <span className="font-medium text-slate-800">{formatCurrency(total)}</span>
              </div>
            </div>
          </div>

          <div className="flex justify-end gap-4">
            <Button type="button" variant="outline" onClick={() => router.push("/invoices")}>
              Cancelar
            </Button>
            <Button type="submit" className="flex items-center gap-2">
              <LockIcon size={16} />
              Processar Pagamento
            </Button>
          </div>
        </form>
      </CardContent>
    </Card>
  )
}
