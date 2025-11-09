import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Role } from '../_models/role.model';
import { RoleSectionPermission } from '../_models/role-section-permission.model';
import { environment } from '../../environments/environment';

export interface RoleWithPermissions {
  id: string;
  name: string;
  permissions: { [section: string]: string[] };
}

@Injectable({
  providedIn: 'root'
})
export class RoleService {

  private apiUrl = `${environment.apiUrl}/role`;

  constructor(private http: HttpClient) { }

  // Role CRUD
  getAllRoles(): Observable<Role[]> {
    return this.http.get<Role[]>(this.apiUrl);
  }

  getRoleById(id: string): Observable<Role> {
    return this.http.get<Role>(`${this.apiUrl}/${id}`);
  }

  getRoleWithPermissions(id: string): Observable<RoleWithPermissions> {
    return this.http.get<RoleWithPermissions>(`${this.apiUrl}/${id}/permissions`);
  }

  createRole(role: Partial<Role>): Observable<Role> {
    return this.http.post<Role>(this.apiUrl, role);
  }

  updateRole(id: string, role: Partial<Role>): Observable<Role> {
    return this.http.patch<Role>(`${this.apiUrl}/${id}`, role);
  }

  deleteRole(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`);
  }

  // Permission management
  getRolePermissions(roleId: string): Observable<RoleSectionPermission[]> {
    return this.http.get<RoleSectionPermission[]>(`${this.apiUrl}/${roleId}/permissions`);
  }

  assignPermissionToRole(assignment: {
    roleId: string;
    sectionId: string;
    permissionId: string;
  }): Observable<RoleSectionPermission> {
    return this.http.post<RoleSectionPermission>(`${this.apiUrl}/permissions`, assignment);
  }

  removePermissionFromRole(roleId: string, sectionId: string, permissionId: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${roleId}/permissions/${sectionId}/${permissionId}`);
  }

  // Backward compatibility
  getRoles(): Observable<Role[]> {
    return this.getAllRoles();
  }
}
