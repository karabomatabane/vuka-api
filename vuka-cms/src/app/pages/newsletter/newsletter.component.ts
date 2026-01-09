import { CommonModule } from '@angular/common';
import {
  Component,
  AfterViewInit,
  ViewChild,
  ChangeDetectionStrategy,
  OnInit,
  inject,
  ChangeDetectorRef,
} from '@angular/core';

import { finalize } from 'rxjs/operators';

import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { Router, RouterLink } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { NewsletterSubscriber } from 'src/app/_models/newsletter-subscriber.model';
import { NewsletterService } from 'src/app/_services/newsletter.service';

@Component({
  selector: 'app-newsletter',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatSortModule,
    MatPaginatorModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatTooltipModule,
    MatButtonModule,
    MatSnackBarModule,
    RouterLink,
  ],
  templateUrl: './newsletter.component.html',
  styleUrls: ['./newsletter.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NewsletterComponent implements OnInit, AfterViewInit {
  private newsletterService = inject(NewsletterService);
  private cdr = inject(ChangeDetectorRef);
  private snackBar = inject(MatSnackBar);

  displayedColumns: string[] = ['name', 'email', 'phone', 'createdAt', 'actions'];
  dataSource = new MatTableDataSource<NewsletterSubscriber>([]);
  isLoading = true;

  @ViewChild(MatSort) sort!: MatSort;
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.loadSubscribers();
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
    this.dataSource.paginator = this.paginator;
  }

  loadSubscribers() {
    this.isLoading = true;
    this.newsletterService
      .getSubscribers()
      .pipe(
        finalize(() => {
          this.isLoading = false;
          this.cdr.detectChanges();
        })
      )
      .subscribe({
        next: (data) => {
          this.dataSource.data = data;
        },
        error: (err) => {
          console.error('Error fetching subscribers:', err);
          this.snackBar.open('Error loading subscribers', 'Close', {
            duration: 3000,
          });
        },
      });
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

  deleteSubscriber(subscriber: NewsletterSubscriber) {
    if (confirm(`Are you sure you want to delete ${subscriber.preferredName}?`)) {
      this.newsletterService.deleteSubscriber(subscriber.id).subscribe(() => {
        this.loadSubscribers();
        this.snackBar.open('Subscriber deleted', 'Close', {
          duration: 3000,
        });
      });
    }
  }
}