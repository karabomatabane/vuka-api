import { BaseModel } from './base.model';
import { DirectoryEntry } from './directory-entry.model';

export interface DirectoryCategory extends BaseModel {
  name: string;
  directories: DirectoryEntry[];
}
