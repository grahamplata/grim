package text

type Sol struct{}

// There are 10^9 lamports in one SOL
const LAMPORTS_PER_SOL uint64 = 1000000000

// LamportsToSol approximately convert fractional native tokens (lamports) into native tokens SOL
func LamportsToSol(lamports uint64) uint64 {
	return lamports / LAMPORTS_PER_SOL
}

// SolToLamports approximately converts native tokens SOL into fractional native tokens
func SolToLamports(sol uint64) uint64 {
	return sol * LAMPORTS_PER_SOL
}
