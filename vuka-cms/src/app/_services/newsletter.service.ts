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

  getTemplate() {
    return this.http.get<{ template: string }>(`${this.baseUrl}/template`);
  }

  updateTemplate(template: string) {
    return this.http.put(`${this.baseUrl}/template`, { template });
  }

  previewNewsletter(articleLimit: number = 5, templateData?: any) {
    return this.http.post(`${this.baseUrl}/preview`, 
      { articleLimit, templateData },
      { responseType: 'text' }
    );
  }

  sendNewsletter(subject: string, content: string, useTemplate: boolean = false, templateData?: any) {
    return this.http.post(`${this.baseUrl}/send`, {
      subject,
      content,
      useTemplate,
      templateData
    });
  }

  sendNewsletterWithArticles(subject: string, limit: number = 5) {
    return this.http.post(`${this.baseUrl}/send/articles`, {
      subject,
      limit
    });
  }

  sendTestEmail(email: string, name: string) {
    return this.http.post(`${this.baseUrl}/test-email`, { email, name });
  }

  testSMTPConnection() {
    return this.http.get(`${this.baseUrl}/test-smtp`);
  }
}
