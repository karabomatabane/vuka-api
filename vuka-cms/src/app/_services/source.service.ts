import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Source } from '../_models/source.model';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SourceService {
  private readonly baseUrl = `${environment.apiUrl}/source`;

  constructor(private http: HttpClient) { }

  getSources() {
    return this.http.get<Source[]>(this.baseUrl);
  }

  getSourceById(id: string) {
    return this.http.get<Source>(`${this.baseUrl}/${id}`);
  }

  createSource(source: Partial<Source>) {
    return this.http.post<Source>(this.baseUrl, source);
  }

  updateSource(id: string, source: Partial<Source>) {
    return this.http.put<Source>(`${this.baseUrl}/${id}`, source);
  }

  deleteSource(id: string) {
    return this.http.delete(`${this.baseUrl}/${id}`);
  }

  ingestSourceFeed(sourceId: string) {
    return this.http.post(`${this.baseUrl}/${sourceId}/ingest`, {});
  }
}
