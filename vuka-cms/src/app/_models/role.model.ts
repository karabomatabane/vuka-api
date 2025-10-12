import { BaseModel } from './base.model';
import { RoleSectionPermission } from './role-section-permission.model';

export interface Role extends BaseModel {
  name: string;
  roleSectionPermissions: RoleSectionPermission[];
}
