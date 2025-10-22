import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { DirectoryCategory } from '../_models/directory-category.model';
import { DirectoryEntry } from '../_models/directory-entry.model';
import { Dir } from '@angular/cdk/bidi';

@Injectable({
  providedIn: 'root',
})
export class DirectoryService {
  private apiUrl = 'http://localhost:3000/directory';

  constructor(private http: HttpClient) {}

  getDirectories() {
    return this.http.get<any[]>(this.apiUrl);
  }

  getDirectoryEntries(categoryId: string) {
    return this.http.get<DirectoryEntry[]>(
      `${this.apiUrl}/entries/${categoryId}`
    );
  }

  createDirectoryCategory(category: DirectoryCategory) {
    return this.http.post<any>(this.apiUrl, category);
  }

  createDirectoryEntry(entry: DirectoryEntry) {
    return this.http.post<DirectoryEntry>(`${this.apiUrl}/entries`, entry);
  }

  updateDirectoryEntry(entryId: string, entry: DirectoryEntry) {
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
