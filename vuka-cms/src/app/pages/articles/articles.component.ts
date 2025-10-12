import {
  Component,
  AfterViewInit,
  ViewChild,
  ChangeDetectionStrategy,
  OnInit,
  inject,
} from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { HttpClientModule } from '@angular/common/http'; // <-- Import HttpClientModule
import { finalize } from 'rxjs/operators';

// Import Angular Material modules
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner'; // <-- For loading indicator
import { MatTooltipModule } from '@angular/material/tooltip';
import { ArticleService } from 'src/app/_services/article.service';
import { Article } from 'src/app/_models/article.model';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { Router } from '@angular/router';

@Component({
  selector: 'app-articles',
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
    DatePipe,
  ],
  providers: [ArticleService], // <-- Provide the service to the component
  templateUrl: './articles.component.html',
  styleUrl: './articles.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ArticlesComponent implements OnInit, AfterViewInit {
  private articleService = inject(ArticleService);
  private router = inject(Router);

  displayedColumns: string[] = [
    'isFeatured',
    'title',
    'sourceName',
    'publishedAt',
  ];
  dataSource = new MatTableDataSource<Article>([]);
  isLoading = true; // Start in a loading state

  @ViewChild(MatSort) sort!: MatSort;
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.loadArticles();
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
    this.dataSource.paginator = this.paginator;
    
    this.dataSource.sortingDataAccessor = (item, property) => {
      switch (property) {
        case 'sourceName':
          return item.source?.name || '';
        default:
          return (item as any)[property];
      }
    };
  }

  loadArticles() {
    this.isLoading = true;
    this.articleService
      .getArticles()
      .pipe(
        finalize(() => (this.isLoading = false)) // Ensure loading is turned off on complete or error
      )
      .subscribe({
        next: (data) => {
          this.dataSource.data = data as Article[];
        },
        error: (err) => {
          console.error('Error fetching articles:', err);
          // Optionally, display an error message to the user
        },
      });
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage(); // Go back to the first page on filter
    }
  }

  openArticleDetails(article: Article) {
    this.router.navigate(['/articles', article.id]);
  }
}
