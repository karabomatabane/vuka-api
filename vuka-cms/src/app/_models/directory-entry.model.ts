import { BaseModel } from './base.model';
import { DirectoryCategory } from './directory-category.model';

export interface DirectoryEntry extends BaseModel {
  name: string;
  description: string;
  contactInfo: ContactInfo[];
  websiteUrl: string;
  entryType: string;
  categoryId: string;
  category: DirectoryCategory;
}

export interface ContactInfo extends BaseModel {
  type: ContactType;
  description: string;
  value: string;
  directoryEntryId?: string;
}

export type ContactType = 'phone' | 'email' | 'fax' | 'address' | 'linkedin' | 'twitter' | 'other';

export const CONTACT_TYPES: { value: ContactType; label: string }[] = [
  { value: 'email', label: 'Email' },
  { value: 'phone', label: 'Phone' },
  { value: 'address', label: 'Address' },
  { value: 'fax', label: 'Fax' },
  { value: 'linkedin', label: 'LinkedIn' },
  { value: 'twitter', label: 'Twitter' },
  { value: 'other', label: 'Other' },
];
