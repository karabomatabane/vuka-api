import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {
  private apiUrl = 'http://localhost:3000/article';
  constructor(private http: HttpClient) { }

  getArticles() {
    return this.http.get(this.apiUrl);
  }

  getArticleById(id: string) {
    return this.http.get(`${this.apiUrl}/${id}`);
  }

  createArticle(article: any) {
    return this.http.post(this.apiUrl, article);
  }

  updateArticle(id: string, article: any) {
    return this.http.put(`${this.apiUrl}/${id}`, article);
  }

  deleteArticle(id: number) {
    return this.http.delete(`${this.apiUrl}/${id}`);
  }
}
