import { Permission } from './permission.model';
import { Role } from './role.model';
import { Section } from './section.model';

export interface RoleSectionPermission {
  roleId: string;
  sectionId: string;
  permissionId: string;
  role: Role;
  section: Section;
  permission: Permission;
}
