import { Injectable, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

import { User } from '../_models/user.model';
import { AuthResponse } from '../_models/auth.model';
import { environment } from 'src/environments/environment';

@Injectable({ providedIn: 'root' })
export class AuthenticationService {
  currentUser = signal<AuthResponse | null>(JSON.parse(localStorage.getItem('currentUser') || 'null'));
  private readonly baseUrl = environment.apiUrl + '/auth';

  constructor(private http: HttpClient) { }

  login(username: string, password: string) {
    return this.http.post<AuthResponse>(`${this.baseUrl}/login`, { username, password })
      .pipe(map(user => {
        localStorage.setItem('currentUser', JSON.stringify(user));
        this.currentUser.set(user);
        return user;
      }));
  }

  logout() {
    localStorage.removeItem('currentUser');
    this.currentUser.set(null);
  }
}
