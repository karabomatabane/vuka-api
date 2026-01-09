import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';
import { PaginatedArticles } from '../_models/article.model';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {
  private readonly baseUrl = `${environment.apiUrl}/article`;
  constructor(private http: HttpClient) { }

  getArticles(page: number, pageSize: number, search?: string): Observable<PaginatedArticles> {
    let url = `${this.baseUrl}?page=${page}&pageSize=${pageSize}`;
    if (search) {
      url += `&search=${search}`;
    }
    return this.http.get<PaginatedArticles>(url);
  }

  getArticleById(id: string) {
    return this.http.get(`${this.baseUrl}/${id}`);
  }

  createArticle(article: any) {
    return this.http.post(this.baseUrl, article);
  }

  updateArticle(id: string, article: any) {
    return this.http.put(`${this.baseUrl}/${id}`, article);
  }

  deleteArticle(id: number) {
    return this.http.delete(`${this.baseUrl}/${id}`);
  }
}
