import { BaseModel } from './base.model';
import { DirectoryCategory } from './directory-category.model';

export interface DirectoryEntry extends BaseModel {
  name: string;
  description: string;
  contactInfo: string;
  websiteUrl: string;
  entryType: string;
  categoryId: string;
  category: DirectoryCategory;
}
