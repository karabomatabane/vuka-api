import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Source } from '../_models/source.model';

@Injectable({
  providedIn: 'root'
})
export class SourceService {
  private apiUrl = 'http://localhost:3000/source';

  constructor(private http: HttpClient) { }

  getSources() {
    return this.http.get<Source[]>(this.apiUrl);
  }

  getSourceById(id: string) {
    return this.http.get<Source>(`${this.apiUrl}/${id}`);
  }

  createSource(source: Partial<Source>) {
    return this.http.post<Source>(this.apiUrl, source);
  }

  updateSource(id: string, source: Partial<Source>) {
    return this.http.put<Source>(`${this.apiUrl}/${id}`, source);
  }

  deleteSource(id: string) {
    return this.http.delete(`${this.apiUrl}/${id}`);
  }

  ingestSourceFeed(sourceId: string) {
    return this.http.post(`${this.apiUrl}/${sourceId}/ingest`, {});
  }
}
