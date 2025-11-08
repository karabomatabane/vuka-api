import { Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthenticationService } from 'src/app/_services/auth.service';
import { first } from 'rxjs/operators';

import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatIconModule } from '@angular/material/icon';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatIconModule,
    MatSnackBarModule
],
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
  private fb = inject(FormBuilder);
  private router = inject(Router);
  private authService = inject(AuthenticationService);
  private snackBar = inject(MatSnackBar);

  loginForm = this.fb.group({
    username: ['', Validators.required],
    password: ['', Validators.required],
  });

  loading = false;
  submitted = false;
  snackBarConfig = {
    duration: 3000,
    panelClass: 'snack-bar-container',
  };

  hide = signal(true);
  clickEvent(event: MouseEvent) {
    this.hide.set(!this.hide());
    event.stopPropagation();
  }

  onSubmit() {
    this.submitted = true;

    if (this.loginForm.invalid) {
      return;
    }

    this.loading = true;
    this.authService
      .login(this.loginForm.value.username!, this.loginForm.value.password!)
      .pipe(first())
      .subscribe({
        next: () => {
          this.snackBar.open('Welcome back!', 'Close', this.snackBarConfig);
          this.router.navigate(['/']);
        },
        error: (error) => {
          this.snackBar.open(error.error.error, 'Close', this.snackBarConfig);
          this.loading = false;
        },
      });
  }

  showRegisterMessage() {
    this.snackBar.open(
      'Registration is disabled. Please contact the administrator.',
      'Close',
      this.snackBarConfig
    );
  }
}
