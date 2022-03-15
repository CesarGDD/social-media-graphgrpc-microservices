import { Comment, Post, User } from "./src/lib/types";

export const users:User[] = [
    {
        id: 1,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    },
    {
        id: 2,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    },
    {
        id: 3,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    },
    {
        id: 4,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    },
    {
        id: 5,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    },
    {
        id: 6,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    },
    {
        id: 7,
        created_at: 12,
        updated_at: 12,
        username: "Carnicero",
        bio: "some fake bio",
        avatar: "/photo.jpeg",
        email: "test@test",
    }
]

export const posts:Post[] = [
    {
        id: 1,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 1,
        user: {
                id: 1,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 2,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 1,
        user: {
                id: 1,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 3,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 1,
        user: {
                id: 1,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 4,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 1,
        user: {
                id: 1,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 5,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 2,
        user: {
                id: 2,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 6,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 2,
        user: {
                id: 2,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 7,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 2,
        user: {
                id: 2,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 8,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 3,
        user: {
                id: 3,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 9,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 3,
        user: {
                id: 3,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 10,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 3,
        user: {
                id: 3,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
    {
        id: 11,
        created_at: 13,
        updated_at: 14,
        url: "/photo.jpeg",
        caption: "Some caption",
        user_id: 3,
        user: {
                id: 3,
                created_at: 12,
                updated_at: 12,
                username: "Carnicero",
                bio: "some fake bio",
                avatar: "/photo.jpeg",
                email: "test@test",
            },
    },
]

export const comments:Comment[]=[
    {
        id: 1,
        created_at: 10,
        updated_at: 10,
        contents: "Que pedo puto",
        user_id: 1,
        post_id: 1,
        user: {
            id: 1,
            created_at: 12,
            updated_at: 12,
            username: "Carnicero",
            bio: "some fake bio",
            avatar: "CA",
            email: "test@test",
        }
    },
    {
        id: 1,
        created_at: 10,
        updated_at: 10,
        contents: "Que pedo puto",
        user_id: 2,
        post_id: 1,
        user: {
            id: 2,
            created_at: 12,
            updated_at: 12,
            username: "Carnicero",
            bio: "some fake bio",
            avatar: "CA",
            email: "test@test",
        }
    },
    {
        id: 1,
        created_at: 10,
        updated_at: 10,
        contents: "Que pedo puto",
        user_id: 2,
        post_id: 1,
        user: {
            id: 2,
            created_at: 12,
            updated_at: 12,
            username: "Carnicero",
            bio: "some fake bio",
            avatar: "CA",
            email: "test@test",
        }
    },
    {
        id: 1,
        created_at: 10,
        updated_at: 10,
        contents: "Que pedo puto",
        user_id: 1,
        post_id: 2,
        user: {
            id: 1,
            created_at: 12,
            updated_at: 12,
            username: "Carnicero",
            bio: "some fake bio",
            avatar: "CA",
            email: "test@test",
        }
    },
    {
        id: 1,
        created_at: 10,
        updated_at: 10,
        contents: "Que pedo puto",
        user_id: 1,
        post_id: 2,
        user: {
            id: 1,
            created_at: 12,
            updated_at: 12,
            username: "Carnicero",
            bio: "some fake bio",
            avatar: "CA",
            email: "test@test",
        }
    },
    {
        id: 1,
        created_at: 10,
        updated_at: 10,
        contents: "Que pedo puto",
        user_id: 1,
        post_id: 2,
        user: {
            id: 1,
            created_at: 12,
            updated_at: 12,
            username: "Carnicero",
            bio: "some fake bio",
            avatar: "CA",
            email: "test@test",
        }
    },
]