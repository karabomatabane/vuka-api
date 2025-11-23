import {
  Component,
  AfterViewInit,
  ViewChild,
  ChangeDetectionStrategy,
  OnInit,
  inject,
} from '@angular/core';
import { CommonModule } from '@angular/common';
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
import { Article, PaginatedArticles } from 'src/app/_models/article.model';
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
  ],
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
  totalArticles = 0;

  @ViewChild(MatSort) sort!: MatSort;
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.loadArticles();
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
    this.paginator.page.subscribe(() => this.loadArticles());
  }

  loadArticles() {
    this.isLoading = true;
    this.articleService
      .getArticles(this.paginator?.pageIndex ?? 1, this.paginator?.pageSize ?? 10)
      .pipe(
        finalize(() => (this.isLoading = false)) // Ensure loading is turned off on complete or error
      )
      .subscribe({
        next: (data: PaginatedArticles) => {
          this.dataSource.data = data.data;
          this.totalArticles = data.pagination.total;
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
