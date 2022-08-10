export interface Post {
    id: number;
    title: string;
    markup: string;
    category: string;
    upvotes: number;
    createdAt: Date;
    updatedAt: Date;
}