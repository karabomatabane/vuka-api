import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { DirectoryCategory, DirectoryOverview } from '../_models/directory-category.model';
import { DirectoryEntry } from '../_models/directory-entry.model';
import { Dir } from '@angular/cdk/bidi';
import { O } from '@angular/cdk/overlay-module.d-B3qEQtts';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class DirectoryService {
  private apiUrl = 'http://localhost:3000/directory';

  constructor(private http: HttpClient) {}

  getOverview(): Observable<DirectoryOverview> {
    return this.http.get<DirectoryOverview>(`${this.apiUrl}/overview`);
  }

  getDirectories(): Observable<DirectoryCategory[]> {
    return this.http.get<DirectoryCategory[]>(this.apiUrl);
  }

  getDirectoryEntries(categoryId: string): Observable<DirectoryEntry[]> {
    return this.http.get<DirectoryEntry[]>(
      `${this.apiUrl}/entries/${categoryId}`
    );
  }

  createDirectoryCategory(category: Partial<DirectoryCategory>) {
    return this.http.post<any>(this.apiUrl, category);
  }

  createDirectoryEntry(entry: Partial<DirectoryEntry>) {
    return this.http.post<DirectoryEntry>(`${this.apiUrl}/entries`, entry);
  }

  updateDirectoryEntry(entryId: string, entry: Partial<DirectoryEntry>) {
    return this.http.put<DirectoryEntry>(
      `${this.apiUrl}/entries/${entryId}`,
      entry
    );
  }

  deleteDirectoryEntry(entryId: string) {
    return this.http.delete<DirectoryEntry>(
      `${this.apiUrl}/entries/${entryId}`
    );
  }
}
