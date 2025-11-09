import { Component, OnInit } from '@angular/core';
import { UserService } from '../../_services/user.service';
import { RoleService, RoleWithPermissions } from '../../_services/role.service';
import { PermissionService } from '../../_services/permission.service';
import { User } from '../../_models/user.model';
import { Role } from '../../_models/role.model';
import { Permission } from '../../_models/permission.model';
import { Section } from '../../_models/section.model';

import { FormsModule } from '@angular/forms';
import { MatTabsModule } from '@angular/material/tabs';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatChipsModule } from '@angular/material/chips';

@Component({
  selector: 'app-roles-and-permissions',
  standalone: true,
  imports: [
    FormsModule,
    MatTabsModule,
    MatTableModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatChipsModule
],
  templateUrl: './roles-and-permissions.component.html',
  styleUrl: './roles-and-permissions.component.scss',
})
export class RolesAndPermissionsComponent implements OnInit {
  // Users
  users: User[] = [];
  selectedUser?: User | null;
  isCreatingUser = false;
  newUser: Partial<User> = {};
  
  // Roles
  roles: Role[] = [];
  selectedRole?: Role | null;
  isCreatingRole = false;
  newRole: Partial<Role> = {};
  selectedRolePermissions?: RoleWithPermissions;
  
  // Permissions
  permissions: Permission[] = [];
  selectedPermission?: Permission | null;
  isCreatingPermission = false;
  newPermission: Partial<Permission> = {};
  
  // Sections
  sections: Section[] = [];
  selectedSection?: Section | null;
  isCreatingSection = false;
  newSection: Partial<Section> = {};

  // For role permission assignment
  isAssigningPermission = false;
  permissionAssignment = {
    roleId: '',
    sectionId: '',
    permissionId: ''
  };

  constructor(
    private userService: UserService,
    private roleService: RoleService,
    private permissionService: PermissionService
  ) {}

  ngOnInit(): void {
    this.loadData();
  }

  loadData(): void {
    this.loadUsers();
    this.loadRoles();
    this.loadPermissions();
    this.loadSections();
  }

  cancelAllEdits(): void {
    this.selectedUser = null;
    this.selectedRole = null;
    this.selectedPermission = null;
    this.selectedSection = null;
    this.isCreatingUser = false;
    this.isCreatingRole = false;
    this.isCreatingPermission = false;
    this.isCreatingSection = false;
    this.isAssigningPermission = false;
    this.newUser = {};
    this.newRole = {};
    this.newPermission = {};
    this.newSection = {};
    this.selectedRolePermissions = undefined;
  }

  // User Management
  loadUsers(): void {
    this.userService.getAllUsers().subscribe(users => this.users = users);
  }

  selectUser(user: User): void {
    this.selectedUser = { ...user };
    this.isCreatingUser = false;
  }

  startCreatingUser(): void {
    this.isCreatingUser = true;
    this.selectedUser = null;
    this.newUser = { username: '', password: '', roleId: '' };
  }

  createUser(): void {
    if (!this.newUser.username || !this.newUser.password || !this.newUser.roleId) return;
    
    this.userService.createUser(this.newUser).subscribe(() => {
      this.loadUsers();
      this.cancelAllEdits();
    });
  }

  updateUser(): void {
    if (!this.selectedUser) return;
    
    this.userService.updateUser(this.selectedUser.id, {
      username: this.selectedUser.username,
      roleId: this.selectedUser.roleId
    }).subscribe(() => {
      this.loadUsers();
      this.cancelAllEdits();
    });
  }

  deleteUser(userId: string): void {
    if (!confirm('Are you sure you want to delete this user?')) return;
    
    this.userService.deleteUser(userId).subscribe(() => {
      this.loadUsers();
      this.cancelAllEdits();
    });
  }

  // Role Management
  loadRoles(): void {
    this.roleService.getAllRoles().subscribe(roles => this.roles = roles);
  }

  selectRole(role: Role): void {
    this.selectedRole = { ...role };
    this.isCreatingRole = false;
    this.loadRolePermissions(role.id);
  }

  loadRolePermissions(roleId: string): void {
    this.roleService.getRoleWithPermissions(roleId).subscribe(
      rolePerms => this.selectedRolePermissions = rolePerms
    );
  }

  startCreatingRole(): void {
    this.isCreatingRole = true;
    this.selectedRole = null;
    this.newRole = { name: '' };
  }

  createRole(): void {
    if (!this.newRole.name) return;
    
    this.roleService.createRole(this.newRole).subscribe(() => {
      this.loadRoles();
      this.cancelAllEdits();
    });
  }

  updateRole(): void {
    if (!this.selectedRole) return;
    
    this.roleService.updateRole(this.selectedRole.id, {
      name: this.selectedRole.name
    }).subscribe(() => {
      this.loadRoles();
      this.cancelAllEdits();
    });
  }

  deleteRole(roleId: string): void {
    if (!confirm('Are you sure you want to delete this role?')) return;
    
    this.roleService.deleteRole(roleId).subscribe(() => {
      this.loadRoles();
      this.cancelAllEdits();
    });
  }

  // Role Permission Assignment
  startAssigningPermission(roleId: string): void {
    this.isAssigningPermission = true;
    this.permissionAssignment = { roleId, sectionId: '', permissionId: '' };
  }

  assignPermission(): void {
    if (!this.permissionAssignment.roleId || 
        !this.permissionAssignment.sectionId || 
        !this.permissionAssignment.permissionId) return;
    
    this.roleService.assignPermissionToRole(this.permissionAssignment).subscribe(() => {
      if (this.selectedRole) {
        this.loadRolePermissions(this.selectedRole.id);
      }
      this.isAssigningPermission = false;
      this.permissionAssignment = { roleId: '', sectionId: '', permissionId: '' };
    });
  }

  removePermission(roleId: string, sectionId: string, permissionId: string): void {
    if (!confirm('Are you sure you want to remove this permission?')) return;
    
    this.roleService.removePermissionFromRole(roleId, sectionId, permissionId).subscribe(() => {
      if (this.selectedRole) {
        this.loadRolePermissions(this.selectedRole.id);
      }
    });
  }

  // Permission Management
  loadPermissions(): void {
    this.permissionService.getAllPermissions().subscribe(perms => this.permissions = perms);
  }

  selectPermission(permission: Permission): void {
    this.selectedPermission = { ...permission };
    this.isCreatingPermission = false;
  }

  startCreatingPermission(): void {
    this.isCreatingPermission = true;
    this.selectedPermission = null;
    this.newPermission = { name: '' };
  }

  createPermission(): void {
    if (!this.newPermission.name) return;
    
    this.permissionService.createPermission(this.newPermission).subscribe(() => {
      this.loadPermissions();
      this.cancelAllEdits();
    });
  }

  updatePermission(): void {
    if (!this.selectedPermission) return;
    
    this.permissionService.updatePermission(this.selectedPermission.id, {
      name: this.selectedPermission.name
    }).subscribe(() => {
      this.loadPermissions();
      this.cancelAllEdits();
    });
  }

  deletePermission(permissionId: string): void {
    if (!confirm('Are you sure you want to delete this permission?')) return;
    
    this.permissionService.deletePermission(permissionId).subscribe(() => {
      this.loadPermissions();
      this.cancelAllEdits();
    });
  }

  // Section Management
  loadSections(): void {
    this.permissionService.getAllSections().subscribe(sects => this.sections = sects);
  }

  selectSection(section: Section): void {
    this.selectedSection = { ...section };
    this.isCreatingSection = false;
  }

  startCreatingSection(): void {
    this.isCreatingSection = true;
    this.selectedSection = null;
    this.newSection = { name: '' };
  }

  createSection(): void {
    if (!this.newSection.name) return;
    
    this.permissionService.createSection(this.newSection).subscribe(() => {
      this.loadSections();
      this.cancelAllEdits();
    });
  }

  updateSection(): void {
    if (!this.selectedSection) return;
    
    this.permissionService.updateSection(this.selectedSection.id, {
      name: this.selectedSection.name
    }).subscribe(() => {
      this.loadSections();
      this.cancelAllEdits();
    });
  }

  deleteSection(sectionId: string): void {
    if (!confirm('Are you sure you want to delete this section?')) return;
    
    this.permissionService.deleteSection(sectionId).subscribe(() => {
      this.loadSections();
      this.cancelAllEdits();
    });
  }

  // Helper methods
  getPermissionsAsArray(permissions?: { [section: string]: string[] }): { section: string, perms: string[] }[] {
    if (!permissions) return [];
    return Object.entries(permissions).map(([section, perms]) => ({ section, perms }));
  }
}
