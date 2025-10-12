import { BaseModel } from './base.model';
import { Role } from './role.model';

export interface User extends BaseModel {
  username: string;
  password?: string;
  roleId: string;
  role: Role;
}
