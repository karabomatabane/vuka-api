import { ChangeDetectorRef, Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { SourceService } from 'src/app/_services/source.service';
import { Source } from 'src/app/_models/source.model';
import { finalize } from 'rxjs';

@Component({
  selector: 'app-source-edit',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './source-edit.component.html',
  styleUrls: ['./source-edit.component.scss'],
})
export class SourceEditComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private fb = inject(FormBuilder);
  private sourceService = inject(SourceService);
  private cdr = inject(ChangeDetectorRef);

  source: Source | undefined;
  editForm!: FormGroup;
  isLoading = true;
  isNew = false;

  ngOnInit() {
    const sourceId = this.route.snapshot.paramMap.get('id');
    this.isNew = !sourceId;

    this.editForm = this.fb.group({
      name: ['', Validators.required],
      websiteUrl: ['', Validators.required],
      rssFeedUrl: ['', Validators.required],
    });

    if (sourceId) {
      this.sourceService
        .getSourceById(sourceId)
        .pipe(
          finalize(() => {
            this.isLoading = false;
            this.cdr.detectChanges();
          })
        )
        .subscribe({
          next: (data) => {
            this.source = data;
            this.editForm.patchValue(this.source);
          },
          error: (err) => {
            console.error('Error fetching source:', err);
          },
        });
    } else {
      this.isLoading = false;
    }
  }

  save() {
    if (this.editForm.valid) {
      this.isLoading = true;
      const formValue = this.editForm.value;
      const observable = this.isNew
        ? this.sourceService.createSource(formValue)
        : this.sourceService.updateSource(this.source!.id, formValue);

      observable
        .pipe(
          finalize(() => {
            this.isLoading = false;
            this.cdr.detectChanges();
          })
        )
        .subscribe({
          next: () => {
            this.router.navigate(['/sources']);
          },
          error: (err) => {
            console.error('Error saving source:', err);
          },
        });
    }
  }

  cancel() {
    this.router.navigate(['/sources']);
  }
}
