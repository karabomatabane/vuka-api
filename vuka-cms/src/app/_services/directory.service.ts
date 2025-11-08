import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { DirectoryCategory, DirectoryOverview } from '../_models/directory-category.model';
import { DirectoryEntry } from '../_models/directory-entry.model';
import { Dir } from '@angular/cdk/bidi';
import { O } from '@angular/cdk/overlay-module.d-B3qEQtts';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class DirectoryService {
  private readonly baseUrl = `${environment.apiUrl}/directory`;

  constructor(private http: HttpClient) {}

  getOverview(): Observable<DirectoryOverview> {
    return this.http.get<DirectoryOverview>(`${this.baseUrl}/overview`);
  }

  getDirectories(): Observable<DirectoryCategory[]> {
    return this.http.get<DirectoryCategory[]>(this.baseUrl);
  }

  getDirectoryEntries(categoryId: string): Observable<DirectoryCategory> {
    return this.http.get<DirectoryCategory>(
      `${this.baseUrl}/entries/${categoryId}`
    );
  }

  getEntryById(entryId: string): Observable<DirectoryEntry> {
    return this.http.get<DirectoryEntry>(`${this.baseUrl}/entry/${entryId}`);
  }

  createDirectoryCategory(category: Partial<DirectoryCategory>) {
    return this.http.post<any>(this.baseUrl, category);
  }

  createDirectoryEntry(entry: Partial<DirectoryEntry>) {
    return this.http.post<DirectoryEntry>(`${this.baseUrl}/entries`, entry);
  }

  updateDirectoryEntry(entryId: string, entry: Partial<DirectoryEntry>) {
    return this.http.put<DirectoryEntry>(
      `${this.baseUrl}/entries/${entryId}`,
      entry
    );
  }

  deleteDirectoryEntry(entryId: string) {
    return this.http.delete<DirectoryEntry>(
      `${this.baseUrl}/entries/${entryId}`
    );
  }
}
