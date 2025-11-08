import { D } from '@angular/cdk/bidi-module.d-D-fEBKdS';

import { Component, inject, signal } from '@angular/core';
import {
  FormBuilder,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import {
  MAT_DIALOG_DATA,
  MatDialogRef,
  MatDialogContent,
  MatDialogActions,
  MatDialogModule,
} from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import {
  MatProgressSpinnerModule,
} from '@angular/material/progress-spinner';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { first, min } from 'rxjs';
import { DirectoryCategory } from 'src/app/_models/directory-category.model';
import { AuthenticationService } from 'src/app/_services/auth.service';
import { DirectoryService } from 'src/app/_services/directory.service';

@Component({
  selector: 'app-directory-form-dialog',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatProgressSpinnerModule,
    MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatDialogModule,
    MatButtonModule,
    MatDialogContent,
    MatDialogActions
],
  templateUrl: './directory-form-dialog.component.html',
  styleUrl: './directory-form-dialog.component.scss',
})
export class DirectoryFormDialogComponent {
  readonly dialogRef = inject(MatDialogRef<DirectoryFormDialogComponent>);
  readonly data = inject(MAT_DIALOG_DATA);
  private fb = inject(FormBuilder);
  private router = inject(Router);
  private directoryService = inject(DirectoryService);
  private snackBar = inject(MatSnackBar);

  directoryForm = this.fb.group({
    name: ['', [Validators.required, Validators.minLength(3)]],
  });

  loading = false;
  submitted = false;
  snackBarConfig = {
    duration: 3000,
    panelClass: 'snack-bar-container',
  };

  onSubmit() {
    this.submitted = true;

    if (this.directoryForm.invalid) {
      return;
    }

    this.loading = true;
    const directoryCategory = {
      name: this.directoryForm.value.name!,
    };
    this.directoryService
      .createDirectoryCategory(directoryCategory)
      .pipe(first())
      .subscribe({
        next: () => {
          this.snackBar.open(
            'Directory created!',
            'Close',
            this.snackBarConfig,
          );
          this.router.navigate(['/']);
          this.dialogRef.close(true);
        },
        error: (error) => {
          this.snackBar.open(error.error.error, 'Close', this.snackBarConfig);
          this.loading = false;
          this.dialogRef.close(false);
        },
      });
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
