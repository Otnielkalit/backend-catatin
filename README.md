# Backend Catatain API Documentation

Dokumentasi endpoint untuk aplikasi Backend Catatain. Semua endpoint mengembalikan standard response format:

```json
{
    "message": "success",
    "status": 200,
    "data": { ... }
}
```

**Base URL**: `https://welcoming-radiance-production-4303.up.railway.app/api/v1`

---

## 1. Users

### Login / Register
Mendaftar user baru (jika belum ada) atau login jika data cocok.

- **URL**: `/users/login`
- **Method**: `POST`
- **Body (JSON)**:
  ```json
  {
      "username": "otniel",
      "phone_number": "081234567890",
      "pin": "123456"
  }
  ```

---

## 2. Categories

### Create Category
- **URL**: `/categories/`
- **Method**: `POST`
- **Body (JSON)**:
  ```json
  {
      "user_id": 1,
      "name": "Belanja Bulanan"
  }
  ```

### Get All Categories
- **URL**: `/categories/`
- **Method**: `GET`
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

### Get Detail Category
- **URL**: `/categories/:id`
- **Method**: `GET`
- **Path Param**: `id` (Category ID)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

### Delete Category
- **URL**: `/categories/:id`
- **Method**: `DELETE`
- **Path Param**: `id` (Category ID)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

---

## 3. Budgets

### Create Budget
- **URL**: `/budgets/`
- **Method**: `POST`
- **Body (JSON)**:
  ```json
  {
      "user_id": 1,
      "amount": 1500000,
      "month": 3,
      "year": 2026
  }
  ```

### Get All Budgets
- **URL**: `/budgets/`
- **Method**: `GET`
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

### Get Detail Budget
- **URL**: `/budgets/:id`
- **Method**: `GET`
- **Path Param**: `id` (Budget ID)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

### Delete Budget
- **URL**: `/budgets/:id`
- **Method**: `DELETE`
- **Path Param**: `id` (Budget ID)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

---

## 4. Expenses (Transactions)

### Create Expense (With Image)
Untuk membuat pengeluaran baru beserta bukti struk/gambar. Data dikirim menggunakan **Multipart Form-Data**.

- **URL**: `/expenses/`
- **Method**: `POST`
- **Body (Form-Data)**:
  - `user_id`: 1 *(Text)*
  - `category_id`: 2 *(Text)*
  - `title`: "Beli Makan Siang" *(Text)*
  - `amount`: 50000 *(Text)*
  - `transaction_date`: "2026-03-03" *(Text, Format: YYYY-MM-DD)*
  - `image`: [File Gambar] *(File)*

### Get All Expenses
- **URL**: `/expenses/`
- **Method**: `GET`
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

### Get Detail Expense
- **URL**: `/expenses/:id`
- **Method**: `GET`
- **Path Param**: `id` (Expense ID)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

### Delete Expense
Menghapus expense sekaligus akan otomatis menghapus gambar struk yang ada di Cloudinary.

- **URL**: `/expenses/:id`
- **Method**: `DELETE`
- **Path Param**: `id` (Expense ID)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```

---

## 5. Analytics

### Get Expense Analytics
Mendapatkan data analitik pengeluaran berdasarkan bulan dan tahun, yang dikelompokkan berdasarkan kategori. Berguna untuk menampilkan chart di frontend.

Parameter `month` dan `year` dikirim melalui **Query Parameter** (`?month=3&year=2026`), sedangkan `user_id` dikirim melalui **JSON Body**.

- **URL**: `/analytics/expenses`
- **Method**: `GET`
- **Query Params**: `?month=[int]&year=[int]` (Opsional - default melihat semua waktu)
- **Body (JSON)**:
  ```json
  {
      "user_id": 1
  }
  ```
