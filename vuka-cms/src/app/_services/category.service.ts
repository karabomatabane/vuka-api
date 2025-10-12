import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Category } from '../_models/category.model';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private apiUrl = 'http://localhost:3000/category';

  constructor(private http: HttpClient) { }

  getCategories() {
    return this.http.get<Category[]>(this.apiUrl);
  }
}
