import { ChangeDetectorRef, Component, Inject, ViewChild } from '@angular/core';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { ActivatedRoute, ParamMap, Router, RouterLink } from '@angular/router';
import { AnimationItem } from 'lottie-web';
import { AnimationOptions, LottieComponent } from 'ngx-lottie';
import { BaseError } from 'src/app/_models/base-error.model';
import { DirectoryCategory } from 'src/app/_models/directory-category.model';
import { DirectoryService } from 'src/app/_services/directory.service';
import { MatIcon, MatIconModule } from '@angular/material/icon';
import { MatButton, MatButtonModule } from '@angular/material/button';
import { DirectoryEntryFormDialogComponent } from '../../../components/directory-entry-form-dialog/directory-entry-form-dialog.component';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { DirectoryEntry } from 'src/app/_models/directory-entry.model';
import { finalize } from 'rxjs';
import {
  MatFormField,
  MatFormFieldModule,
  MatLabel,
} from '@angular/material/form-field';
import {
  MatProgressSpinner,
  MatProgressSpinnerModule,
} from '@angular/material/progress-spinner';
import { MatInputModule } from '@angular/material/input';
import { MatTooltipModule } from '@angular/material/tooltip';
import { TitleCasePipe } from '@angular/common';

@Component({
  selector: 'app-directory-category',
  standalone: true,
  imports: [
    LottieComponent,
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
    TitleCasePipe,
  ],
  templateUrl: './directory-category.component.html',
  styleUrl: './directory-category.component.scss',
})
export class DirectoryCategoryComponent {
  entryId: string = '';
  directoryCategory: DirectoryCategory | undefined;
  displayedColumns: string[] = [
    'name',
    'description',
    'websiteUrl',
    'entryType',
    'actions',
  ];
  dataSource = new MatTableDataSource<DirectoryEntry>([]);
  isLoading = true;

  @ViewChild(MatSort) sort!: MatSort;
  @ViewChild(MatPaginator) paginator!: MatPaginator;

  snackBarConfig = {
    duration: 3000,
    panelClass: 'snack-bar-container',
  };
  options: AnimationOptions = {
    path: '/lottie/no-data.json',
  };
  dialogConfig: MatDialogConfig = {
    enterAnimationDuration: 300,
    exitAnimationDuration: 300,
  };

  constructor(
    private directoryService: DirectoryService,
    private route: ActivatedRoute,
    private snackBar: MatSnackBar,
    private dialog: MatDialog,
    private cdr: ChangeDetectorRef,
    private router: Router,
  ) {}

  ngOnInit() {
    this.route.paramMap.subscribe((data: ParamMap) => {
      if (this.entryId === data.get('id')) return;
      this.entryId = data.get('id') || '';
      this.loadEntries(this.entryId);
    });
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
    this.dataSource.paginator = this.paginator;
  }

  animationCreated(animationItem: AnimationItem): void {
    console.log(animationItem);
  }

  loadEntries(id: string) {
    this.isLoading = true;

    this.directoryService
      .getDirectoryEntries(id)
      .pipe(
        finalize(() => {
          this.isLoading = false;
          this.cdr.detectChanges();
        }),
      )
      .subscribe({
        next: (data: DirectoryCategory) => {
          this.directoryCategory = data;
          this.dataSource.data = data.entries;
        },
        error: (error: BaseError) => {
          console.error('Error fetching directory entries:', error);
          this.snackBar.open(
            'There was an error fetching directory',
            'Close',
            this.snackBarConfig,
          );
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

  onAddEntry() {
    const dialogRef = this.dialog.open(DirectoryEntryFormDialogComponent, {
      ...this.dialogConfig,
      data: {
        categoryName: this.directoryCategory?.name,
      },
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        // Refresh the entries list
        this.loadEntries(this.entryId);
      }
    });
  }
}
