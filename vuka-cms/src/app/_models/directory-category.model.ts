import { BaseModel } from './base.model';
import { DirectoryEntry } from './directory-entry.model';

export interface DirectoryCategory extends BaseModel {
  name: string;
  entries: DirectoryEntry[];
}

export interface OverviewDirectoryCategory extends BaseModel {
  id: string;
  name: string;
  totalEntries: number;
}


export interface DirectoryOverview {
  categories: OverviewDirectoryCategory[];
  personalised: {
    pinned: OverviewDirectoryCategory[];
    recent: OverviewDirectoryCategory[];
  };
}
