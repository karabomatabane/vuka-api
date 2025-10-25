import { BaseModel } from './base.model';
import { User } from './user.model';
import { DirectoryEntry } from './directory-entry.model';

export interface UserDirectoryMeta extends BaseModel {
  userId: string;
  user: User;
  directoryId: string;
  directory: DirectoryEntry;
  pinned: boolean;
  lastAccessed: string;
}
