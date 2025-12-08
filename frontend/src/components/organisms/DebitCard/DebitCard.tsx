import CardComponent from "@/components/molecules/CardComponent";

export default function DebitCard() {
  return (
    <CardComponent title="Debit Card Account">
      {/* Card preview */}
      <div className="bg-linear-to-br from-teal-500 to-teal-700 rounded-xl p-6 text-white mb-4 relative overflow-hidden">
        <div className="absolute top-0 right-0 w-24 h-24 bg-white/10 rounded-full -mr-8 -mt-8" />

        <div className="relative z-10">
          <p className="text-xs opacity-75 mb-8">Focused</p>
          <div className="flex items-center gap-2 mb-6">
            <div className="w-8 h-6 bg-white/20 rounded flex items-center justify-center">
              <span className="text-xs font-bold">💳</span>
            </div>
            <span className="text-sm font-semibold">Platinum Debit</span>
          </div>

          <p className="text-lg font-mono tracking-widest mb-6">
            4771 6080 1080 7897
          </p>

          <div className="flex items-center justify-between">
            <div>
              <p className="text-xs opacity-75">Valid Thru</p>
              <p className="text-sm font-mono">08/25</p>
            </div>
            <div className="w-12 h-8 bg-linear-to-br from-yellow-300 to-orange-400 rounded">
              <span className="text-xs font-bold text-white">VISA</span>
            </div>
          </div>
        </div>
      </div>

      {/* Add card button */}
      <button className="w-full py-3 border-2 border-dashed border-slate-300 dark:border-slate-600 rounded-lg text-sm font-medium text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-700/50 transition-colors">
        + Add Debit Card
      </button>
    </CardComponent>
  );
}
