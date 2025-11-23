import { Component, computed, Input, inject } from '@angular/core';
import { signal } from '@angular/core';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';

import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { MatTooltipModule } from '@angular/material/tooltip';
import { AuthenticationService } from 'src/app/_services/auth.service';
import { DirectoryService } from 'src/app/_services/directory.service';
import { DirectoryCategory } from 'src/app/_models/directory-category.model';
import { MenuItemComponent } from "../menu-item/menu-item.component";

export type MenuItem = {
  icon: string;
  label: string;
  route?: string;
  subItems?: MenuItem[];
};

@Component({
  selector: 'app-x-sidenav',
  standalone: true,
  imports: [
    MatListModule,
    MatIconModule,
    MenuItemComponent
],
  templateUrl: './x-sidenav.component.html',
  styleUrls: ['./x-sidenav.component.scss'],
})
export class XSidenavComponent {
  private authService = inject(AuthenticationService);
  private directoryService = inject(DirectoryService);
  private router = inject(Router);

  currentUser = this.authService.currentUser;
  directoryCategories = signal<DirectoryCategory[]>([]);

  ngOnInit() {
    this.directoryService.getDirectories().subscribe((categories) => {
      this.directoryCategories.set(categories);
    });
  }

  sideNavCollapsed = signal(false);
  subMenuOpen = signal(false);
  @Input() set collapsed(val: boolean) {
    this.sideNavCollapsed.set(val);
  }

  profilePicSize = computed(() => (this.sideNavCollapsed() ? '32' : '100'));

  menuItems = computed(() => {
    if (this.currentUser()) {
      return [
        { icon: 'newsstand', label: 'Articles', route: '/articles' },
        { icon: 'home_storage', label: 'Sources', route: '/sources' },
        {
          icon: 'inventory_2',
          label: 'Directories',
          route: '/directories',
          subItems: [
            ...this.directoryCategories().map((category) => ({
              icon: 'folder',
              label: category.name,
              route: `/directories/${category.id}`,
            })),
          ],
        },
        { icon: 'mail', label: 'Newsletter', route: '/newsletter' },
        {
          icon: 'security',
          label: 'Roles & Permissions',
          route: '/roles-and-permissions',
        },
      ];
    } else {
      return [{ icon: 'login', label: 'Login', route: '/login' }];
    }
  });

  logoutItem: MenuItem = { icon: 'logout', label: 'Logout' };

  logout = () => {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
}
