import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { NewsletterSubscriber } from '../_models/newsletter-subscriber.model';

@Injectable({
  providedIn: 'root'
})
export class NewsletterService {
  private readonly baseUrl = `${environment.apiUrl}/newsletter`;
  constructor(private http: HttpClient) { }

  getSubscribers() {
    return this.http.get<NewsletterSubscriber[]>(this.baseUrl+'/subscribers');
  }

  deleteSubscriber(id: string) {
    return this.http.delete(`${this.baseUrl}/subscribers/${id}`);
  }
}
