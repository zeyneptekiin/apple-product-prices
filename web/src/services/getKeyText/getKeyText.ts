import countries from '../../json_files/countries.json';

type Countries = typeof countries;
export type CountryCode = keyof Countries;

export function getKeyText(code: CountryCode): string {
    return countries[code] || 'Unknown Country';
}
