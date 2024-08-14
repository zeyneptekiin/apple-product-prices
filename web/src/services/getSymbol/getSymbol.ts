import symbols from '../../json_files/symbol.json';

type Symbols = typeof symbols;
export type CurrencySymbol = keyof Symbols;

export function getSymbol(code: CurrencySymbol): string {
    return symbols[code] || 'Unknown Symbol';
}
