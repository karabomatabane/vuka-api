import {
  Component,
  AfterViewInit,
  ViewChild,
  ChangeDetectionStrategy,
  OnInit,
  inject,
  ChangeDetectorRef,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { finalize } from 'rxjs/operators';

// Import Angular Material modules
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { Router, RouterLink } from '@angular/router';
import { Source } from 'src/app/_models/source.model';
import { SourceService } from 'src/app/_services/source.service';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';

@Component({
  selector: 'app-sources',
  standalone: true,
  imports: [
    CommonModule,
    RouterLink,
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
  ],
  providers: [SourceService],
  templateUrl: './sources.component.html',
  styleUrls: ['./sources.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SourcesComponent implements OnInit, AfterViewInit {
  private sourceService = inject(SourceService);
  private router = inject(Router);
  private cdr = inject(ChangeDetectorRef);
  private snackBar = inject(MatSnackBar);

  displayedColumns: string[] = ['name', 'websiteUrl', 'rssFeedUrl', 'actions'];
  dataSource = new MatTableDataSource<Source>([]);
  isLoading = true;

  @ViewChild(MatSort) sort!: MatSort;
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  ngOnInit() {
    this.loadSources();
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
    this.dataSource.paginator = this.paginator;
  }

  loadSources() {
    this.isLoading = true;
    this.sourceService
      .getSources()
      .pipe(finalize(() => {
        this.isLoading = false;
        this.cdr.detectChanges();
      }))
      .subscribe({
        next: (data) => {
          this.dataSource.data = data;
        },
        error: (err) => {
          console.error('Error fetching sources:', err);
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

  editSource(source: Source) {
    this.router.navigate(['/sources', source.id, 'edit']);
  }

  deleteSource(source: Source) {
    if (confirm(`Are you sure you want to delete ${source.name}?`)) {
      this.sourceService.deleteSource(source.id).subscribe(() => {
        this.loadSources();
      });
    }
  }

  ingestFeed(source: Source) {
    this.sourceService.ingestSourceFeed(source.id).subscribe({
      next: (res: any) => {
        this.snackBar.open(res.message, 'Close', { duration: 3000 });
      },
      error: (err) => {
        this.snackBar.open('Error starting feed ingestion', 'Close', { duration: 3000 });
      },
    });
  }
}