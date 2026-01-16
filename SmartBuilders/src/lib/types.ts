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

export type Project = {
  id: string;
  manager_id: string;
  customer_lastname: string;
  address: string;
  description: string;
  start_date: string;
  end_date: string;
  created_at?: string;
};