import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Permission } from '../_models/permission.model';
import { Section } from '../_models/section.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PermissionService {

  private permissionApiUrl = `${environment.apiUrl}/permission`;
  private sectionApiUrl = `${environment.apiUrl}/section`;

  constructor(private http: HttpClient) { }

  // Permission CRUD
  getAllPermissions(): Observable<Permission[]> {
    return this.http.get<Permission[]>(this.permissionApiUrl);
  }

  getPermissionById(id: string): Observable<Permission> {
    return this.http.get<Permission>(`${this.permissionApiUrl}/${id}`);
  }

  createPermission(permission: Partial<Permission>): Observable<Permission> {
    return this.http.post<Permission>(this.permissionApiUrl, permission);
  }

  updatePermission(id: string, permission: Partial<Permission>): Observable<Permission> {
    return this.http.patch<Permission>(`${this.permissionApiUrl}/${id}`, permission);
  }

  deletePermission(id: string): Observable<void> {
    return this.http.delete<void>(`${this.permissionApiUrl}/${id}`);
  }

  // Section CRUD
  getAllSections(): Observable<Section[]> {
    return this.http.get<Section[]>(this.sectionApiUrl);
  }

  getSectionById(id: string): Observable<Section> {
    return this.http.get<Section>(`${this.sectionApiUrl}/${id}`);
  }

  createSection(section: Partial<Section>): Observable<Section> {
    return this.http.post<Section>(this.sectionApiUrl, section);
  }

  updateSection(id: string, section: Partial<Section>): Observable<Section> {
    return this.http.patch<Section>(`${this.sectionApiUrl}/${id}`, section);
  }

  deleteSection(id: string): Observable<void> {
    return this.http.delete<void>(`${this.sectionApiUrl}/${id}`);
  }
}
