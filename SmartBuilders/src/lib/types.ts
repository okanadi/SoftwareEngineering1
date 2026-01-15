export type Field = {
  label: string;
  name: string;
  type?:    
    | 'text'
    | 'email'
    | 'password'
    | 'select'
    | 'textarea'
    | 'date'
    | 'anrede'
    | 'vorname'
    | 'nachname';
  options?: string[];
  conditional?: (values: Record<string, any>) => boolean;
  placeholder?: string;
};