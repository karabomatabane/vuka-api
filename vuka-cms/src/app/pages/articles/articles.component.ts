import {
  Component,
  AfterViewInit,
  ViewChild,
  ChangeDetectionStrategy,
  OnInit,
  inject,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { finalize, Subject, debounceTime } from 'rxjs';

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
  private searchSubject = new Subject<string>();

  displayedColumns: string[] = [
    'isFeatured',
    'title',
    'sourceName',
    'publishedAt',
  ];
  dataSource = new MatTableDataSource<Article>([]);
  isLoading = true; // Start in a loading state
  totalArticles = 0;
  search: string = '';

  @ViewChild(MatSort) sort!: MatSort;
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.searchSubject.pipe(
      debounceTime(300)
    ).subscribe(searchValue => {
      this.search = searchValue;
      this.paginator.pageIndex = 0;
      this.loadArticles();
    });
  }

  ngAfterViewInit() {
    this.sort.sortChange.subscribe(() => this.loadArticles());
    this.paginator.page.subscribe(() => this.loadArticles());
    this.loadArticles();
  }

  loadArticles() {
    this.isLoading = true;
    this.articleService
      .getArticles(
        this.paginator?.pageIndex + 1,
        this.paginator?.pageSize,
        this.search
      )
      .pipe(
        finalize(() => (this.isLoading = false))
      )
      .subscribe({
        next: (data: PaginatedArticles) => {
          this.dataSource.data = data.data;
          this.totalArticles = data.pagination.totalItems;
        },
        error: (err) => {
          console.error('Error fetching articles:', err);
        },
      });
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.searchSubject.next(filterValue.trim().toLowerCase());
  }

  openArticleDetails(article: Article) {
    this.router.navigate(['/articles', article.id]);
  }
}
