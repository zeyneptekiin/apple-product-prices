import countries from '../../json_files/currency.json';

type Countries = typeof countries;
export type CountryCode = keyof Countries;

export function getCurrency(code: CountryCode): string {
    return countries[code] || 'Unknown Country';
}
