import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  FormsModule,
  ReactiveFormsModule,
  Validators,
  FormArray,
  FormGroup,
} from '@angular/forms';
import {
  MatDialogContent,
  MatDialogActions,
  MAT_DIALOG_DATA,
  MatDialogRef,
  MatDialogModule,
} from '@angular/material/dialog';
import {
  MatFormFieldModule,
} from '@angular/material/form-field';
import {
  MatProgressSpinnerModule,
} from '@angular/material/progress-spinner';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { first } from 'rxjs';
import { BaseError } from 'src/app/_models/base-error.model';
import { DirectoryService } from 'src/app/_services/directory.service';
import { DirectoryFormDialogComponent } from '../directory-form-dialog/directory-form-dialog.component';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatIconModule } from '@angular/material/icon';
import { ContactType, CONTACT_TYPES, DirectoryEntry } from 'src/app/_models/directory-entry.model';

@Component({
  selector: 'app-directory-entry-form-dialog',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatProgressSpinnerModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatIconModule,
    FormsModule,
    MatDialogModule,
    MatButtonModule,
    MatDialogContent,
    MatDialogActions,
  ],
  templateUrl: './directory-entry-form-dialog.component.html',
  styleUrl: './directory-entry-form-dialog.component.scss',
})
export class DirectoryEntryFormDialogComponent {
  readonly dialogRef = inject(MatDialogRef<DirectoryFormDialogComponent>);
  readonly data = inject(MAT_DIALOG_DATA);
  private fb = inject(FormBuilder);
  private router = inject(Router);
  private directoryService = inject(DirectoryService);
  private snackBar = inject(MatSnackBar);
  
  categoryName = this.data.categoryName || 'Unknown';
  categoryId = this.data.categoryId;
  contactTypes = CONTACT_TYPES;

  entryForm = this.fb.group({
    name: ['', [Validators.required, Validators.minLength(3)]],
    description: ['', [Validators.required, Validators.maxLength(216)]],
    websiteUrl: [''],
    entryType: ['business'],
    categoryId: [this.categoryId, Validators.required],
    contactInfo: this.fb.array([this.createContactInfoGroup()])
  });

  loading = false;
  submitted = false;
  snackBarConfig = {
    duration: 3000,
    panelClass: 'snack-bar-container',
  };

  get contactInfoArray(): FormArray {
    return this.entryForm.get('contactInfo') as FormArray;
  }

  createContactInfoGroup(): FormGroup {
    return this.fb.group({
      type: ['email' as ContactType, Validators.required],
      description: [''],
      value: ['', [Validators.required, this.contactValueValidator.bind(this)]]
    });
  }

  contactValueValidator(control: any) {
    if (!control.value) return null;
    
    const parent = control.parent;
    if (!parent) return null;
    
    const type = parent.get('type')?.value;
    const value = control.value;

    switch (type) {
      case 'email':
        const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailPattern.test(value) ? null : { invalidEmail: true };
      case 'phone':
        // Basic phone validation - allows various formats
        const phonePattern = /^\+[1-9]\d{7,14}$/;
        return phonePattern.test(value) ? null : { invalidPhone: true };
      default:
        return null;
    }
  }

  addContactInfo(): void {
    this.contactInfoArray.push(this.createContactInfoGroup());
  }

  removeContactInfo(index: number): void {
    if (this.contactInfoArray.length > 1) {
      this.contactInfoArray.removeAt(index);
    }
  }

  getPlaceholderForType(type: ContactType): string {
    switch (type) {
      case 'email':
        return 'example@company.com';
      case 'phone':
        return '+27823010000';
      case 'address':
        return '123 Main St, City, State, ZIP';
      case 'fax':
        return '+27821098343';
      case 'linkedin':
        return 'https://linkedin.com/company/example';
      case 'twitter':
        return '@companyhandle';
      default:
        return '';
    }
  }

  onSubmit() {
    this.submitted = true;

    if (this.entryForm.invalid) {
      return;
    }

    this.loading = true;
    const directoryEntry: Partial<DirectoryEntry> = {
      name: this.entryForm.value.name!,
      description: this.entryForm.value.description!,
      websiteUrl: this.entryForm.value.websiteUrl || '',
      entryType: this.entryForm.value.entryType || 'business',
      categoryId: this.entryForm.value.categoryId!,
      contactInfo: this.entryForm.value.contactInfo!
    };

    this.directoryService
      .createDirectoryEntry(directoryEntry)
      .pipe(first())
      .subscribe({
        next: () => {
          this.snackBar.open(
            'Directory entry created!',
            'Close',
            this.snackBarConfig,
          );
          this.dialogRef.close(true);
        },
        error: (error: BaseError) => {
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
