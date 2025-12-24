"use client";

import Link from "next/link";

export default function OwnerDashboardPage() {
  return (
    <div className="flex min-h-screen bg-gray-100">

      {/* SIDEBAR */}
      <aside className="w-64 bg-indigo-600 text-white flex flex-col">
        <div className="p-6 text-2xl font-bold border-b border-indigo-500">
          👑 Owner
        </div>

        <nav className="flex-1 p-4 space-y-3">
          <a className="block px-4 py-2 rounded-lg bg-indigo-700">
            Dashboard
          </a>
          <a className="block px-4 py-2 rounded-lg hover:bg-indigo-700 transition">
            Data Karyawan
          </a>
          <a className="block px-4 py-2 rounded-lg hover:bg-indigo-700 transition">
            Laporan Keuangan
          </a>
          <a className="block px-4 py-2 rounded-lg hover:bg-indigo-700 transition">
            Pengaturan
          </a>
        </nav>

        <div className="p-4 border-t border-indigo-500 text-sm">
          © 2025 Company
        </div>
      </aside>

      {/* MAIN CONTENT */}
      <main className="flex-1 flex flex-col">

        {/* TOP BAR */}
        <header className="flex items-center justify-between bg-white px-6 py-4 shadow">
          <h1 className="text-xl font-semibold text-gray-800">
            Dashboard Owner
          </h1>

          <Link
            href="/register"
            className="text-indigo-600 font-medium hover:underline"
          >
            + Register
          </Link>
        </header>

        {/* CONTENT */}
        <section className="p-6 grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="bg-white p-6 rounded-xl shadow">
            <h2 className="text-gray-500 text-sm">Total Karyawan</h2>
            <p className="text-3xl font-bold mt-2">25</p>
          </div>

          <div className="bg-white p-6 rounded-xl shadow">
            <h2 className="text-gray-500 text-sm">Total Pendapatan</h2>
            <p className="text-3xl font-bold mt-2">Rp 120.000.000</p>
          </div>

          <div className="bg-white p-6 rounded-xl shadow">
            <h2 className="text-gray-500 text-sm">Laporan Bulan Ini</h2>
            <p className="text-3xl font-bold mt-2">12</p>
          </div>
        </section>

      </main>
    </div>
  );
}

