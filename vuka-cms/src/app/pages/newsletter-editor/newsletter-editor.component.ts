import { Component, OnInit, inject, signal, ViewChild, ElementRef, AfterViewChecked } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { MatSliderModule } from '@angular/material/slider';
import { MatDividerModule } from '@angular/material/divider';
import { NewsletterService } from 'src/app/_services/newsletter.service';

@Component({
  selector: 'app-newsletter-editor',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatTabsModule,
    MatFormFieldModule,
    MatInputModule,
    MatProgressSpinnerModule,
    MatSnackBarModule,
    MatSliderModule,
    MatDividerModule
  ],
  templateUrl: './newsletter-editor.component.html',
  styleUrls: ['./newsletter-editor.component.scss']
})
export class NewsletterEditorComponent implements OnInit, AfterViewChecked {
  private newsletterService = inject(NewsletterService);
  private snackBar = inject(MatSnackBar);

  @ViewChild('previewFrame') previewFrame?: ElementRef<HTMLIFrameElement>;

  templateContent = signal<string>('');
  previewHtml = signal<string>('');
  isLoading = signal<boolean>(false);
  isSaving = signal<boolean>(false);
  isPreviewing = signal<boolean>(false);
  articleLimit = signal<number>(5);
  selectedTab = signal<number>(0);
  private needsIframeUpdate = false;

  // Send newsletter
  sendSubject = '';
  isSending = false;

  ngOnInit() {
    this.loadTemplate();
  }

  ngAfterViewChecked() {
    if (this.needsIframeUpdate && this.previewFrame?.nativeElement) {
      // Use setTimeout to ensure iframe is ready
      setTimeout(() => this.updateIframeContent(), 0);
      this.needsIframeUpdate = false;
    }
  }

  private updateIframeContent() {
    const iframe = this.previewFrame?.nativeElement;
    if (!iframe) return;

    const doc = iframe.contentDocument || iframe.contentWindow?.document;
    if (doc) {
      doc.open();
      doc.write(this.previewHtml());
      doc.close();
    }
  }

  loadTemplate() {
    this.isLoading.set(true);
    this.newsletterService.getTemplate().subscribe({
      next: (response) => {
        this.templateContent.set(response.template);
        this.isLoading.set(false);
        this.generatePreview();
      },
      error: (err) => {
        console.error('Error loading template:', err);
        this.snackBar.open('Failed to load template', 'Close', { duration: 3000 });
        this.isLoading.set(false);
      }
    });
  }

  saveTemplate() {
    this.isSaving.set(true);
    this.newsletterService.updateTemplate(this.templateContent()).subscribe({
      next: () => {
        this.snackBar.open('Template saved successfully', 'Close', { duration: 3000 });
        this.isSaving.set(false);
        this.generatePreview();
      },
      error: (err) => {
        console.error('Error saving template:', err);
        this.snackBar.open('Failed to save template', 'Close', { duration: 3000 });
        this.isSaving.set(false);
      }
    });
  }

  generatePreview() {
    this.isPreviewing.set(true);
    this.newsletterService.previewNewsletter(this.articleLimit()).subscribe({
      next: (html) => {
        console.log('Preview HTML received, length:', html.length);
        this.previewHtml.set(html);
        this.needsIframeUpdate = true;
        this.isPreviewing.set(false);
        
        // Force update if already on preview tab
        if (this.selectedTab() === 1) {
          setTimeout(() => this.updateIframeContent(), 100);
        }
      },
      error: (err) => {
        console.error('Error generating preview:', err);
        this.snackBar.open('Failed to generate preview', 'Close', { duration: 3000 });
        this.isPreviewing.set(false);
      }
    });
  }

  onArticleLimitChange(event: any) {
    this.articleLimit.set(event.target.value);
    this.generatePreview();
  }

  onTabChange(event: any) {
    // When switching to preview tab, update iframe
    if (event.index === 1 && this.previewHtml()) {
      setTimeout(() => this.updateIframeContent(), 100);
    }
  }

  sendNewsletter() {
    if (!this.sendSubject.trim()) {
      this.snackBar.open('Please enter a subject', 'Close', { duration: 3000 });
      return;
    }

    this.isSending = true;
    this.newsletterService.sendNewsletterWithArticles(this.sendSubject, this.articleLimit()).subscribe({
      next: () => {
        this.snackBar.open('Newsletter sent successfully!', 'Close', { duration: 5000 });
        this.isSending = false;
        this.sendSubject = '';
      },
      error: (err) => {
        console.error('Error sending newsletter:', err);
        this.snackBar.open('Failed to send newsletter', 'Close', { duration: 3000 });
        this.isSending = false;
      }
    });
  }

  sendTestEmail() {
    const email = prompt('Enter email address for test:');
    if (!email) return;

    const name = prompt('Enter recipient name:') || 'Test User';

    this.newsletterService.sendTestEmail(email, name).subscribe({
      next: () => {
        this.snackBar.open('Test email sent successfully!', 'Close', { duration: 3000 });
      },
      error: (err) => {
        console.error('Error sending test email:', err);
        this.snackBar.open('Failed to send test email', 'Close', { duration: 3000 });
      }
    });
  }

  resetTemplate() {
    if (confirm('Are you sure you want to reset the template? This will reload the saved version.')) {
      this.loadTemplate();
    }
  }
}
